---
page: blog blog_post
codapi: true
authors:
  - Jesús Espino
date: 2026-09-07
enableComments: true
title: "PostgreSQL 19 interactive tour"
summary: "An interactive tour of what's new in PostgreSQL 19: property graph queries, REPACK, plan advice, temporal updates, and more, with short runnable SQL examples you can edit and run in the browser."
categories:
  - PostgreSQL
tags:
  - postgresql
  - postgres
  - sql
  - databases
images:
  - /blog/postgres-19/preview.webp
---

PostgreSQL 19 is in beta, with general availability expected around September or October 2026, so it's a good time to get a head start on what's new. The [official release notes](https://www.postgresql.org/docs/19/release-19.html) are the authoritative record: meticulously assembled, complete down to the commit and the contributors behind each change. This article is a hands-on companion to them, taking a selection of those entries and turning each into a runnable example so you can see how the new behavior actually works.

Before we start digging into the new features, let's set the context.

This article is based on the official release notes and the PostgreSQL source code, licensed under the [PostgreSQL License](https://www.postgresql.org/about/licence/). This is not an exhaustive list; see the [official release notes](https://www.postgresql.org/docs/19/release-19.html) for that.

Every example below was run against **PostgreSQL 19 beta 3** (released 2026-08-13) and the output is what that server actually printed.

Links point to the documentation (𝗗), the most relevant commits (𝗖), and authors (𝗔) for each feature; check them out for motivation, usage, and implementation details. The authors (𝗔) are the people credited in the release notes for the feature, which usually means the patch authors rather than a single main author.

With the context set, let's start exploring the new features.

## Property graph queries (reverted after beta 3)

> **Heads-up:** [reverted](https://github.com/postgres/postgres/commit/b1f106c80) on 2026-09-07, three and a half weeks after beta 3. It won't ship in 19; earliest is 20. The examples still run on beta 3, so treat them as a preview.

This was going to be the headline of the release. PostgreSQL 19 implements [SQL/PGQ](https://www.postgresql.org/docs/19/queries-graph.html), the property-graph part of SQL:2023. You declare a *property graph* over existing tables, then query it with pattern matching instead of writing the joins yourself.

Two ordinary tables, one graph on top of them:

```sql
CREATE PROPERTY GRAPH social
    VERTEX TABLES (
        person KEY (id) LABEL person PROPERTIES (id, name)
    )
    EDGE TABLES (
        follows KEY (follower, followee)
            SOURCE KEY (follower) REFERENCES person (id)
            DESTINATION KEY (followee) REFERENCES person (id)
            LABEL follows
    );
```

Nothing is copied: `social` is a view-like object that says "`person` rows are vertices, `follows` rows are edges". Now you can match patterns with `GRAPH_TABLE`, where `-[...]->` is a directed edge:

```sql
SELECT * FROM GRAPH_TABLE (social
    MATCH (a IS person)-[IS follows]->(b IS person)
    COLUMNS (a.name AS follower, b.name AS followee)
) ORDER BY follower, followee;
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="graph.sql" output></codapi-snippet>

```text
┌──────────┬──────────┐
│ follower │ followee │
├──────────┼──────────┤
│ Ada      │ Bo       │
│ Ada      │ Dee      │
│ Bo       │ Cleo     │
│ Cleo     │ Dee      │
└──────────┴──────────┘
(4 rows)
```

The payoff is multi-hop patterns. Chaining two edges gives you friends-of-friends without a self-join, and an empty `()` means "some vertex I don't care to name":

```sql
SELECT * FROM GRAPH_TABLE (social
    MATCH (a IS person WHERE a.name = 'Ada')-[IS follows]->()-[IS follows]->(c IS person)
    COLUMNS (a.name AS start, c.name AS friend_of_friend)
);
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="graph.sql" output></codapi-snippet>

```text
┌───────┬──────────────────┐
│ start │ friend_of_friend │
├───────┼──────────────────┤
│ Ada   │ Cleo             │
└───────┴──────────────────┘
(1 row)
```

There's no new execution engine here, and that's the point. `GRAPH_TABLE` is rewritten into a plain relational query, so the planner, the statistics, and the index choices you already know all still apply:

```sql
EXPLAIN (COSTS OFF)
SELECT * FROM GRAPH_TABLE (social
    MATCH (a IS person)-[IS follows]->(b IS person)
    COLUMNS (a.name AS follower, b.name AS followee)
);
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="graph.sql" output></codapi-snippet>

```text
┌───────────────────────────────────────────────────┐
│                    QUERY PLAN                     │
├───────────────────────────────────────────────────┤
│ Hash Join                                         │
│   Hash Cond: (follows.followee = person_1.id)     │
│   ->  Hash Join                                   │
│         Hash Cond: (follows.follower = person.id) │
│         ->  Seq Scan on follows                   │
│         ->  Hash                                  │
│               ->  Seq Scan on person              │
│   ->  Hash                                        │
│         ->  Seq Scan on person person_1           │
└───────────────────────────────────────────────────┘
(9 rows)
```

One limitation to know before you plan a migration off a graph database: this first cut has no variable-length paths. Quantifiers like `-[IS follows]->{1,3}` parse but are rejected with `element pattern quantifier is not supported`, so a pattern has to spell out every hop.

- 𝗗 [Graph Queries](https://www.postgresql.org/docs/19/queries-graph.html), [Property Graphs](https://www.postgresql.org/docs/19/ddl-property-graphs.html), [CREATE PROPERTY GRAPH](https://www.postgresql.org/docs/19/sql-create-property-graph.html)
- 𝗖 [2f094e7](https://github.com/postgres/postgres/commit/2f094e7ac), [c5b3253](https://github.com/postgres/postgres/commit/c5b3253b8), [a0dd070](https://github.com/postgres/postgres/commit/a0dd0702e), reverted by [b1f106c](https://github.com/postgres/postgres/commit/b1f106c80)
- 𝗔 Peter Eisentraut, Ashutosh Bapat

## Temporal updates and deletes

The new [`FOR PORTION OF`](https://www.postgresql.org/docs/19/sql-update.html) clause on `UPDATE` and `DELETE` operates on a slice of a range column. Instead of rewriting a whole validity period, you name a sub-period and PostgreSQL splits the row for you.

The table starts with a single row: one price valid for all of 2026. Change it for July only, and query it again:

```sql
SELECT * FROM price ORDER BY valid_at;

UPDATE price
    FOR PORTION OF valid_at FROM '2026-07-01' TO '2026-08-01'
    SET amount = 7.99;

SELECT * FROM price ORDER BY valid_at;
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="temporal.sql" output></codapi-snippet>

```text
┌────────┬─────────────────────────┬────────┐
│  sku   │        valid_at         │ amount │
├────────┼─────────────────────────┼────────┤
│ widget │ [2026-01-01,2027-01-01) │   9.99 │
└────────┴─────────────────────────┴────────┘
(1 row)

UPDATE 1
┌────────┬─────────────────────────┬────────┐
│  sku   │        valid_at         │ amount │
├────────┼─────────────────────────┼────────┤
│ widget │ [2026-01-01,2026-07-01) │   9.99 │
│ widget │ [2026-07-01,2026-08-01) │   7.99 │
│ widget │ [2026-08-01,2027-01-01) │   9.99 │
└────────┴─────────────────────────┴────────┘
(3 rows)
```

One row in, three rows out, and the untouched periods keep the old price.

`DELETE` works the same way, trimming instead of splitting. This snippet starts from the same untouched full-year row; deleting from December to the end of time leaves the earlier portion behind:

```sql
SELECT * FROM price ORDER BY valid_at;

DELETE FROM price
    FOR PORTION OF valid_at FROM '2026-12-01' TO NULL;

SELECT * FROM price ORDER BY valid_at;
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="temporal.sql" output></codapi-snippet>

```text
┌────────┬─────────────────────────┬────────┐
│  sku   │        valid_at         │ amount │
├────────┼─────────────────────────┼────────┤
│ widget │ [2026-01-01,2027-01-01) │   9.99 │
└────────┴─────────────────────────┴────────┘
(1 row)

DELETE 1
┌────────┬─────────────────────────┬────────┐
│  sku   │        valid_at         │ amount │
├────────┼─────────────────────────┼────────┤
│ widget │ [2026-01-01,2026-12-01) │   9.99 │
└────────┴─────────────────────────┴────────┘
(1 row)
```

`NULL` as a bound means unbounded, so `TO NULL` is "from December onwards". This lands alongside a [new documentation chapter on temporal tables](https://www.postgresql.org/docs/19/ddl-temporal-tables.html), which is worth reading if you keep history in range columns.

- 𝗗 [Temporal Tables](https://www.postgresql.org/docs/19/ddl-temporal-tables.html), [UPDATE](https://www.postgresql.org/docs/19/sql-update.html)
- 𝗖 [8e72d91](https://github.com/postgres/postgres/commit/8e72d914c), [b6ccd30](https://github.com/postgres/postgres/commit/b6ccd30d8)
- 𝗔 Paul A. Jungwirth

## Upsert that returns the row you lost to

`INSERT ... ON CONFLICT DO NOTHING ... RETURNING` has always had an annoying hole: the rows that conflicted are simply absent from the result, so you can't tell "already there" from "never happened". To fix that, you could use `ON CONFLICT DO UPDATE`, but that only works if you are willing to write to every conflicting row.

PostgreSQL 19 adds [`ON CONFLICT DO SELECT`](https://www.postgresql.org/docs/19/sql-insert.html), which returns the *existing* rows without touching them. Here `widget` already exists with `qty = 7`:

```sql
INSERT INTO inventory VALUES ('widget', 1), ('gadget', 3)
    ON CONFLICT (sku) DO SELECT
    RETURNING sku, qty;
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="upsert.sql" output></codapi-snippet>

```text
┌────────┬─────┐
│  sku   │ qty │
├────────┼─────┤
│ widget │   7 │
│ gadget │   3 │
└────────┴─────┘
(2 rows)

INSERT 0 2
```

Both rows come back, and `widget` reports `7` (the value already in the table), not the `1` we tried to insert.

`DO SELECT` also takes a locking clause, so `FOR UPDATE` holds the conflicting rows while you decide what to do with them. It also gives you a way to tell the two kinds of row apart: locking a row stamps its `xmax`, so the [`xmax = 0` trick](https://stackoverflow.com/questions/39058213/differentiate-inserted-and-updated-rows-in-upsert-using-system-columns) tells you which rows were really inserted:

```sql
INSERT INTO inventory VALUES ('widget', 1), ('gadget', 3)
    ON CONFLICT (sku) DO SELECT FOR UPDATE
    RETURNING sku, qty, xmax = 0 AS was_inserted;
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="upsert.sql" output></codapi-snippet>

```text
┌────────┬─────┬──────────────┐
│  sku   │ qty │ was_inserted │
├────────┼─────┼──────────────┤
│ widget │   7 │ f            │
│ gadget │   3 │ t            │
└────────┴─────┴──────────────┘
(2 rows)

INSERT 0 2
```

The locking clause is what makes that work: without one there is no lock to stamp, `xmax` stays `0` on the conflicting row too, and every row claims `was_inserted = t`. Any locking clause will do, `FOR SHARE` included.

- 𝗗 [INSERT](https://www.postgresql.org/docs/19/sql-insert.html)
- 𝗖 [8832709](https://github.com/postgres/postgres/commit/88327092f)
- 𝗔 Andreas Karlsson, Marko Tiikkaja, Viktor Holmberg

## Window functions can skip NULLs

`lead()`, `lag()`, `first_value()`, `last_value()`, and `nth_value()` now accept the SQL-standard [`IGNORE NULLS`](https://www.postgresql.org/docs/19/functions-window.html) clause (and its default counterpart, `RESPECT NULLS`).

The classic use for this is filling gaps in sparse data. Here a sensor only reports when the temperature changes, leaving NULLs in between:

```sql
SELECT ts, temp,
       last_value(temp) IGNORE NULLS OVER (ORDER BY ts) AS carried,
       lag(temp)        IGNORE NULLS OVER (ORDER BY ts) AS prev_reported
FROM readings
ORDER BY ts;
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="window.sql" output></codapi-snippet>

```text
┌────┬──────┬─────────┬───────────────┐
│ ts │ temp │ carried │ prev_reported │
├────┼──────┼─────────┼───────────────┤
│  1 │   20 │      20 │               │
│  2 │      │      20 │            20 │
│  3 │      │      20 │            20 │
│  4 │   23 │      23 │            20 │
│  5 │      │      23 │            23 │
└────┴──────┴─────────┴───────────────┘
(5 rows)
```

`carried` is a last-observation-carried-forward fill, and `prev_reported` is the previous *real* reading rather than the previous row. Before 19 this needed a nested subquery with a grouping trick over `count(temp)`; now it's a keyword.

- 𝗗 [Window Functions](https://www.postgresql.org/docs/19/functions-window.html)
- 𝗖 [25a30bb](https://github.com/postgres/postgres/commit/25a30bbd4)
- 𝗔 Oliver Ford, Tatsuo Ishii

## REPACK

`VACUUM FULL` and `CLUSTER` did almost the same thing (rewrite a table to reclaim space) with two confusing names and no way to avoid an `ACCESS EXCLUSIVE` lock. PostgreSQL 19 unifies them under [`REPACK`](https://www.postgresql.org/docs/19/sql-repack.html).

Take a table with half its rows deleted:

```sql
SELECT pg_size_pretty(pg_table_size('events')) AS size;

REPACK events;

SELECT pg_size_pretty(pg_table_size('events')) AS size;
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="repack.sql" output></codapi-snippet>

```text
┌────────┐
│  size  │
├────────┤
│ 736 kB │
└────────┘
(1 row)

REPACK
┌────────┐
│  size  │
├────────┤
│ 360 kB │
└────────┘
(1 row)
```

The important part is the new `CONCURRENTLY` option, which rebuilds the table without an `ACCESS EXCLUSIVE` lock. It works by decoding the changes that land during the rebuild and replaying them, so reads *and writes* keep running:

```sql
-- No ACCESS EXCLUSIVE lock: readers and writers keep working.
REPACK (CONCURRENTLY) events;
```

That one isn't runnable here: the decoding requires `wal_level` to be `replica` or higher, and the sandbox runs with `minimal`. It also means `CONCURRENTLY` consumes a replication slot; the new `max_repack_replication_slots` setting (default 5) caps how many can run at once.

What `CLUSTER` used to do (physically order the rows by an index) is now spelled `REPACK ... USING INDEX`. Adding `VERBOSE` makes the command report how it rewrote the table: either an index scan, or a sequential scan followed by a sort.

```sql
-- Physically order the rows by an index (the old CLUSTER behavior).
REPACK (VERBOSE) events USING INDEX events_pkey;
```

```text
INFO:  repacking "public.events" using index scan on "events_pkey"
INFO:  "public.events": found 6 removable, 2500 nonremovable row versions in 87 pages
DETAIL:  0 dead row versions cannot be removed yet.
CPU: user: 0.00 s, system: 0.00 s, elapsed: 0.00 s.
REPACK
```

`VACUUM FULL` and `CLUSTER` still work, so nothing breaks; they're just the old spellings now.

- 𝗗 [REPACK](https://www.postgresql.org/docs/19/sql-repack.html)
- 𝗖 [ac58465](https://github.com/postgres/postgres/commit/ac58465e0), [28d534e](https://github.com/postgres/postgres/commit/28d534e2a), [8fb95a8](https://github.com/postgres/postgres/commit/8fb95a8ab), [e76d8c7](https://github.com/postgres/postgres/commit/e76d8c749)
- 𝗔 Antonin Houska, Mihail Nikalayeu, Álvaro Herrera

## Merge and split partitions (reverted after beta 3)

> **Heads-up:** [reverted](https://github.com/postgres/postgres/commit/3e8bcc864) on 2026-08-27, two weeks after beta 3, "due to multiple design issues which are too late to address in this release cycle". It won't ship in 19; earliest is 20. The examples still run on beta 3, so treat them as a preview.

Reshaping a partitioned table used to be a manual dance of `DETACH`, create, `INSERT ... SELECT`, `ATTACH`. The new [`ALTER TABLE ... MERGE PARTITIONS`](https://www.postgresql.org/docs/19/sql-altertable.html) and `ALTER TABLE ... SPLIT PARTITION` commands do it in one statement, moving the rows for you.

```sql
SELECT tableoid::regclass AS partition, * FROM metrics ORDER BY id;

ALTER TABLE metrics MERGE PARTITIONS (metrics_lo, metrics_hi) INTO metrics_all;

SELECT tableoid::regclass AS partition, * FROM metrics ORDER BY id;
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="partitions-merge.sql" output></codapi-snippet>

```text
┌────────────┬────┬───┐
│ partition  │ id │ v │
├────────────┼────┼───┤
│ metrics_lo │  1 │ a │
│ metrics_hi │  5 │ b │
└────────────┴────┴───┘
(2 rows)

ALTER TABLE
┌─────────────┬────┬───┐
│  partition  │ id │ v │
├─────────────┼────┼───┤
│ metrics_all │  1 │ a │
│ metrics_all │  5 │ b │
└─────────────┴────┴───┘
(2 rows)
```

Splitting is the inverse. This example starts from a `metrics` table with a single oversized partition covering the whole range:

```sql
SELECT tableoid::regclass AS partition, * FROM metrics ORDER BY id;

ALTER TABLE metrics SPLIT PARTITION metrics_all INTO (
    PARTITION metrics_lo FOR VALUES FROM (1) TO (4),
    PARTITION metrics_hi FOR VALUES FROM (4) TO (7)
);

SELECT tableoid::regclass AS partition, * FROM metrics ORDER BY id;
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="partitions-split.sql" output></codapi-snippet>

```text
┌─────────────┬────┬───┐
│  partition  │ id │ v │
├─────────────┼────┼───┤
│ metrics_all │  2 │ x │
│ metrics_all │  6 │ y │
└─────────────┴────┴───┘
(2 rows)

ALTER TABLE
┌────────────┬────┬───┐
│ partition  │ id │ v │
├────────────┼────┼───┤
│ metrics_lo │  2 │ x │
│ metrics_hi │  6 │ y │
└────────────┴────┴───┘
(2 rows)
```

The rows land in the right partition automatically, according to the bounds you declared.

Both take an `ACCESS EXCLUSIVE` lock on the parent and rewrite the data, so this is a maintenance-window operation, not an online one. Handy for the common "monthly partitions got too granular, roll them up into quarters" cleanup.

- 𝗗 [ALTER TABLE](https://www.postgresql.org/docs/19/sql-altertable.html)
- 𝗖 [f2e4cc4](https://github.com/postgres/postgres/commit/f2e4cc427), [4b3d173](https://github.com/postgres/postgres/commit/4b3d17362), reverted by [3e8bcc8](https://github.com/postgres/postgres/commit/3e8bcc864)
- 𝗔 Dmitry Koval, Alexander Korotkov, Tender Wang, Richard Guo, Dagfinn Ilmari Mannsåker, Fujii Masao, Jian He

## Autovacuum gets priorities

Autovacuum used to walk tables in whatever order it found them, which meant the table about to cause a wraparound emergency waited behind a dozen boring ones. PostgreSQL 19 gives every table a [score](https://www.postgresql.org/docs/19/routine-vacuuming.html#AUTOVACUUM-PRIORITY) and processes the highest first, and exposes the whole calculation in a new view.

Here's a 20,000-row table with a third of its rows just deleted (rounded for readability, since the columns are `double precision`):

```sql
SELECT relname,
       round(score::numeric, 2)          AS score,
       round(vacuum_score::numeric, 2)   AS vacuum_score,
       round(analyze_score::numeric, 2)  AS analyze_score,
       do_vacuum, do_analyze, for_wraparound
FROM pg_stat_autovacuum_scores
WHERE relname = 'churn';
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="autovac.sql" output></codapi-snippet>

```text
┌─────────┬───────┬──────────────┬───────────────┬───────────┬────────────┬────────────────┐
│ relname │ score │ vacuum_score │ analyze_score │ do_vacuum │ do_analyze │ for_wraparound │
├─────────┼───────┼──────────────┼───────────────┼───────────┼────────────┼────────────────┤
│ churn   │ 13.01 │         1.65 │         13.01 │ t         │ t          │ f              │
└─────────┴───────┴──────────────┴───────────────┴───────────┴────────────┴────────────────┘
(1 row)
```

Read a score as "how far past its threshold this table is". Above 1 means it's due: `vacuum_score` of 1.65 means 65% more dead tuples than it takes to trigger a vacuum. The overall `score` is just the highest of the components, so you can sort by it and see the queue. There are more components in the view (`xid_score`, `mxid_score`, `vacuum_insert_score`), which is what finally makes "why is autovacuum busy on *that* table" answerable.

Each component has a weight you can tune, if freezing matters more than dead tuples in your workload, or set them all to `0.0` to get the pre-19 ordering back:

```sql
SELECT name, setting FROM pg_settings
WHERE name LIKE 'autovacuum%score_weight'
ORDER BY name;
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="autovac.sql" output></codapi-snippet>

```text
┌──────────────────────────────────────────┬─────────┐
│                   name                   │ setting │
├──────────────────────────────────────────┼─────────┤
│ autovacuum_analyze_score_weight          │ 1       │
│ autovacuum_freeze_score_weight           │ 1       │
│ autovacuum_multixact_freeze_score_weight │ 1       │
│ autovacuum_vacuum_insert_score_weight    │ 1       │
│ autovacuum_vacuum_score_weight           │ 1       │
└──────────────────────────────────────────┴─────────┘
(5 rows)
```

- 𝗗 [Autovacuum priority](https://www.postgresql.org/docs/19/routine-vacuuming.html#AUTOVACUUM-PRIORITY), [Autovacuum configuration](https://www.postgresql.org/docs/19/runtime-config-autovacuum.html)
- 𝗖 [d7965d6](https://github.com/postgres/postgres/commit/d7965d65f), [87f61f0](https://github.com/postgres/postgres/commit/87f61f0c8)
- 𝗔 Nathan Bossart, Sami Imseih

## Plan advice, officially

PostgreSQL has refused query hints for its entire history. The new [`pg_plan_advice`](https://www.postgresql.org/docs/19/pgplanadvice.html) module is the compromise: not hints that override the planner, but a way to *record* the plan you have and then constrain the planner to it.

It's a loadable module rather than an extension, so `LOAD` it (or add it to `shared_preload_libraries`) and `EXPLAIN` grows a `PLAN_ADVICE` option:

```sql
EXPLAIN (COSTS OFF, PLAN_ADVICE)
SELECT * FROM fact f JOIN dim d ON f.dim_id = d.id;
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="advice.sql" output></codapi-snippet>

```text
┌────────────────────────────────┐
│           QUERY PLAN           │
├────────────────────────────────┤
│ Hash Join                      │
│   Hash Cond: (f.dim_id = d.id) │
│   ->  Seq Scan on fact f       │
│   ->  Hash                     │
│         ->  Seq Scan on dim d  │
│ Generated Plan Advice:         │
│   JOIN_ORDER(f d)              │
│   HASH_JOIN(d)                 │
│   SEQ_SCAN(f d)                │
│   NO_GATHER(f d)               │
└────────────────────────────────┘
(10 rows)
```

That block is the plan, written in a small advice language: drive from `f`, join `d` to it, hash join with `d` on the inner side, sequential scans, no parallelism. Feed a string like that back through `pg_plan_advice.advice` and the planner is constrained to match. Usually you'd keep only the parts you actually care about:

```sql
SET pg_plan_advice.advice = 'JOIN_ORDER(d f) NESTED_LOOP_PLAIN(f) INDEX_SCAN(f fact_dim_id_idx)';

EXPLAIN (COSTS OFF)
SELECT * FROM fact f JOIN dim d ON f.dim_id = d.id;
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="advice.sql" output></codapi-snippet>

```text
SET
┌──────────────────────────────────────────────────┐
│                    QUERY PLAN                    │
├──────────────────────────────────────────────────┤
│ Nested Loop                                      │
│   ->  Seq Scan on dim d                          │
│   ->  Index Scan using fact_dim_id_idx on fact f │
│         Index Cond: (dim_id = d.id)              │
│ Supplied Plan Advice:                            │
│   INDEX_SCAN(f fact_dim_id_idx) /* matched */    │
│   JOIN_ORDER(d f) /* matched */                  │
│   NESTED_LOOP_PLAIN(f) /* matched */             │
└──────────────────────────────────────────────────┘
(8 rows)
```

The plan flipped to a nested loop driven by `dim`, and every piece of advice reports `/* matched */`. That feedback is the best part of the design: advice that doesn't apply says so instead of silently doing nothing. Ask for something impossible and you get told, plus a plan marked `Disabled: true`:

```sql
-- x is not in the query at all.
SET pg_plan_advice.advice = 'JOIN_ORDER(x f d)';

EXPLAIN (COSTS OFF)
SELECT * FROM fact f JOIN dim d ON f.dim_id = d.id;
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="advice.sql" output></codapi-snippet>

```text
SET
┌──────────────────────────────────────────────────┐
│                    QUERY PLAN                    │
├──────────────────────────────────────────────────┤
│ Nested Loop                                      │
│   Disabled: true                                 │
│   ->  Seq Scan on dim d                          │
│   ->  Index Scan using fact_dim_id_idx on fact f │
│         Index Cond: (dim_id = d.id)              │
│ Supplied Plan Advice:                            │
│   JOIN_ORDER(x f d) /* partially matched */      │
└──────────────────────────────────────────────────┘
(7 rows)
```

Mechanically, advice never *adds* plans; it only removes candidates the planner would otherwise consider. So you can never force a plan the planner considered invalid, and impossible advice degrades to a bad-but-correct plan rather than an error.

The companion [`pg_stash_advice`](https://www.postgresql.org/docs/19/pgstashadvice.html) extension is where this becomes operational: it stores advice keyed by query ID and applies it automatically, so you don't have to `SET` anything from the application.

- 𝗗 [pg_plan_advice](https://www.postgresql.org/docs/19/pgplanadvice.html), [pg_stash_advice](https://www.postgresql.org/docs/19/pgstashadvice.html)
- 𝗖 [5883ff3](https://github.com/postgres/postgres/commit/5883ff30b), [6455e55](https://github.com/postgres/postgres/commit/6455e55b0), [e8ec19a](https://github.com/postgres/postgres/commit/e8ec19aa3), [c10edb1](https://github.com/postgres/postgres/commit/c10edb102)
- 𝗔 Robert Haas, Lukas Fittl

## Aggregate before you join

Also known as eager aggregation. When you group by a column from one side of a join, the planner can now push a partial aggregate *underneath* the join, shrinking the number of rows the join has to process, then finalize on top.

```sql
EXPLAIN (COSTS OFF)
SELECT d.label, count(*)
FROM orders o JOIN dim d ON o.dim_id = d.id
GROUP BY d.label;
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="optimizer.sql" output></codapi-snippet>

```text
┌──────────────────────────────────────────────┐
│                  QUERY PLAN                  │
├──────────────────────────────────────────────┤
│ Finalize GroupAggregate                      │
│   Group Key: d.label                         │
│   ->  Sort                                   │
│         Sort Key: d.label                    │
│         ->  Hash Join                        │
│               Hash Cond: (o.dim_id = d.id)   │
│               ->  Partial HashAggregate      │
│                     Group Key: o.dim_id      │
│                     ->  Seq Scan on orders o │
│               ->  Hash                       │
│                     ->  Seq Scan on dim d    │
└──────────────────────────────────────────────┘
(11 rows)
```

`orders` has 100,000 rows over 100 distinct `dim_id` values. The `Partial HashAggregate` collapses them to 100 rows *before* the join, so the hash join handles 100 rows instead of 100,000. On PostgreSQL 18 the same query aggregates only at the top:

```text
┌──────────────────────────────────────┐
│              QUERY PLAN              │
├──────────────────────────────────────┤
│ HashAggregate                        │
│   Group Key: d.label                 │
│   ->  Hash Join                      │
│         Hash Cond: (o.dim_id = d.id) │
│         ->  Seq Scan on orders o     │
│         ->  Hash                     │
│               ->  Seq Scan on dim d  │
└──────────────────────────────────────┘
(7 rows)
```

This is cost-based, so the planner only does it when the pre-aggregation is expected to pay for itself. The win scales with how much the grouping collapses the input. It has its own switch, `enable_eager_aggregate`, which is `on` by default; set it to `off` and you get the PostgreSQL 18 plan back, which is a quick way to check whether it's helping or hurting a particular query.

- 𝗗 [enable_eager_aggregate](https://www.postgresql.org/docs/19/runtime-config-query.html)
- 𝗖 [8e11859](https://github.com/postgres/postgres/commit/8e1185910), [bd94845](https://github.com/postgres/postgres/commit/bd94845e8), [3a08a2a](https://github.com/postgres/postgres/commit/3a08a2a8b)
- 𝗔 Richard Guo, Antonin Houska

## NOT IN finally becomes an anti-join

`NOT IN (subquery)` has been a performance trap forever. Because SQL's three-valued logic makes `NOT IN` behave strangely if the subquery can produce a NULL, the planner refused to turn it into an anti-join and fell back to a hashed subplan instead.

PostgreSQL 19 proves the NULL case away when it can. With `NOT NULL` on both columns:

```sql
EXPLAIN (COSTS OFF)
SELECT * FROM orders WHERE id NOT IN (SELECT id FROM cancelled);
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="optimizer.sql" output></codapi-snippet>

```text
┌─────────────────────────────────────────┐
│               QUERY PLAN                │
├─────────────────────────────────────────┤
│ Hash Anti Join                          │
│   Hash Cond: (orders.id = cancelled.id) │
│   ->  Seq Scan on orders                │
│   ->  Hash                              │
│         ->  Seq Scan on cancelled       │
└─────────────────────────────────────────┘
(5 rows)
```

PostgreSQL 18 produced this for the same query:

```text
┌──────────────────────────────────────────────────────┐
│                      QUERY PLAN                      │
├──────────────────────────────────────────────────────┤
│ Seq Scan on orders                                   │
│   Filter: (NOT (ANY (id = (hashed SubPlan 1).col1))) │
│   SubPlan 1                                          │
│     ->  Seq Scan on cancelled                        │
└──────────────────────────────────────────────────────┘
(4 rows)
```

The catch is the precondition: **both** the outer expression and the subquery output must be provably non-nullable, which in practice means `NOT NULL` columns. Drop the constraint and you're back to the subplan, which is a decent argument for declaring `NOT NULL` where you can.

Two smaller optimizer changes ship alongside it: more `LEFT JOIN ... WHERE right.col IS NULL` patterns now become anti-joins, and the constant folder simplifies `IS [NOT] DISTINCT FROM` into plain operators when the inputs can't be NULL:

```sql
EXPLAIN (COSTS OFF)
SELECT * FROM cancelled WHERE id IS DISTINCT FROM 42;
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="optimizer.sql" output></codapi-snippet>

```text
┌───────────────────────┐
│      QUERY PLAN       │
├───────────────────────┤
│ Seq Scan on cancelled │
│   Filter: (id <> 42)  │
└───────────────────────┘
(2 rows)
```

On 18 the filter stays as `(id IS DISTINCT FROM 42)`, which can't use an index or a normal operator's statistics. Now it's an ordinary `<>`.

- 𝗗 [Release notes: Optimizer](https://www.postgresql.org/docs/19/release-19.html#RELEASE-19-OPTIMIZER)
- 𝗖 [383eb21](https://github.com/postgres/postgres/commit/383eb21eb), [cf74558](https://github.com/postgres/postgres/commit/cf74558fe), [0a37961](https://github.com/postgres/postgres/commit/0a3796125)
- 𝗔 Richard Guo, Tender Wang

## EXPLAIN knows more

Three additions worth knowing, all in [`EXPLAIN`](https://www.postgresql.org/docs/19/sql-explain.html).

The big one is the new **`IO`** option, which reports what asynchronous I/O actually did. PostgreSQL 18 introduced AIO; now you can see it. Part of the table below is still in shared buffers when the scan runs, so only some of it comes off disk; that part is what the new counters describe:

```sql
EXPLAIN (ANALYZE, IO, COSTS OFF, TIMING OFF, SUMMARY OFF)
SELECT count(*) FROM io_demo;
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="explainio.sql" output></codapi-snippet>

```text
┌─────────────────────────────────────────────────────────────────┐
│                           QUERY PLAN                            │
├─────────────────────────────────────────────────────────────────┤
│ Aggregate (actual rows=1.00 loops=1)                            │
│   Buffers: shared hit=2048 read=320 dirtied=320 written=1       │
│   ->  Seq Scan on io_demo (actual rows=120000.00 loops=1)       │
│         Prefetch: avg=11.35 max=64 capacity=204                 │
│         I/O: count=24 waits=6 size=13.33 in-progress=2.33       │
│         Buffers: shared hit=2048 read=320 dirtied=320 written=1 │
│ Planning:                                                       │
│   Buffers: shared hit=9                                         │
└─────────────────────────────────────────────────────────────────┘
(8 rows)
```

Two new lines. `Prefetch` describes how far ahead the read stream was looking (`avg` and `max` blocks, out of a `capacity` of 204). `I/O` describes the reads themselves: 24 I/O operations averaging 13.33 blocks each, six of which the scan actually had to *wait* for; the rest had already completed by the time it needed them.

This is the one example in this post whose output isn't fixed: `avg`, `waits`, and `in-progress` depend on how fast the kernel returns the reads, and everything depends on how much of the table happens to be cached already, so expect different numbers on your own run. Warm the cache completely and the `I/O` line vanishes, because there was no I/O to report.

`auto_explain` gets the same thing via `auto_explain.log_io`.

Second, **Memoize now explains itself**. A Memoize node used to tell you nothing about why the planner chose it:

```sql
SET enable_hashjoin = off;
SET enable_mergejoin = off;

EXPLAIN SELECT * FROM orders o JOIN dim d ON o.dim_id = d.id;
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="optimizer.sql" output></codapi-snippet>

```text
SET
SET
┌─────────────────────────────────────────────────────────────────────────────────────┐
│                                     QUERY PLAN                                      │
├─────────────────────────────────────────────────────────────────────────────────────┤
│ Nested Loop  (cost=0.15..3944.45 rows=100000 width=15)                              │
│   ->  Seq Scan on orders o  (cost=0.00..1443.00 rows=100000 width=8)                │
│   ->  Memoize  (cost=0.15..0.17 rows=1 width=7)                                     │
│         Cache Key: o.dim_id                                                         │
│         Cache Mode: logical                                                         │
│         Estimates: capacity=100 distinct keys=100 lookups=100000 hit percent=99.90% │
│         ->  Index Scan using dim_pkey on dim d  (cost=0.14..0.16 rows=1 width=7)    │
│               Index Cond: (id = o.dim_id)                                           │
└─────────────────────────────────────────────────────────────────────────────────────┘
(8 rows)
```

That `Estimates:` line is new. It shows the planner expected 100 distinct keys over 100,000 lookups, a 99.9% hit rate, which is exactly why Memoize won. When a Memoize node underperforms in production, you can now put that estimate next to the actual `Hits`/`Misses` that `EXPLAIN ANALYZE` reports and tell whether the planner guessed wrong or the data changed under it.

Third, `EXPLAIN (ANALYZE, WAL)` now separates out how many WAL bytes went to full-page images, as `fpi bytes` next to the existing `fpi` count. Full-page writes are often the bulk of WAL volume right after a checkpoint, and previously you could see the count but not the size.

- 𝗗 [EXPLAIN](https://www.postgresql.org/docs/19/sql-explain.html)
- 𝗖 [681daed](https://github.com/postgres/postgres/commit/681daed93), [3b1117d](https://github.com/postgres/postgres/commit/3b1117d6e), [e157fe6](https://github.com/postgres/postgres/commit/e157fe6f7), [4bc62b8](https://github.com/postgres/postgres/commit/4bc62b868), [5ab0b6a](https://github.com/postgres/postgres/commit/5ab0b6a24)
- 𝗔 Tomas Vondra, Ilia Evdokimov, Lukas Fittl, Shinya Kato

## COPY grows up

Four independent improvements to [`COPY`](https://www.postgresql.org/docs/19/sql-copy.html), and together they cover most of the reasons people reach for a script instead.

`COPY TO` can emit **JSON**, one object per row:

```sql
COPY employee TO STDOUT WITH (FORMAT json);
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="copyx.sql" output></codapi-snippet>

```text
{"id":1,"name":"Ada","dept":"eng"}
{"id":2,"name":"Bo","dept":"ops"}
```

That's JSON Lines, which is what most log and analytics pipelines want. If you need one valid JSON document instead, `FORCE_ARRAY` wraps it:

```sql
COPY employee TO STDOUT WITH (FORMAT json, FORCE_ARRAY true);
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="copyx.sql" output></codapi-snippet>

```text
[
 {"id":1,"name":"Ada","dept":"eng"}
,{"id":2,"name":"Bo","dept":"ops"}
]
```

On the input side, `HEADER` now takes a count, so files with a multi-line preamble no longer need pre-processing, and `ON_ERROR set_null` turns unparseable values into NULLs instead of aborting the whole load. Both at once:

```sql
COPY reading FROM STDIN WITH (FORMAT csv, HEADER 2, ON_ERROR set_null);
id,val
(units: celsius)
1,10
2,oops
3,30
\.

SELECT * FROM reading ORDER BY id;
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="copyx.sql" output></codapi-snippet>

```text
NOTICE:  in 1 row, columns were set to null due to data type incompatibility
COPY 3
┌────┬─────┐
│ id │ val │
├────┼─────┤
│  1 │  10 │
│  2 │     │
│  3 │  30 │
└────┴─────┘
(3 rows)
```

Two header lines skipped, `oops` quietly became NULL, and the NOTICE tells you how many values were affected so the failure isn't silent. This joins `ON_ERROR ignore` from PostgreSQL 17, which dropped the whole row instead of just the bad column.

Fourth: `COPY TO` finally accepts a **partitioned table** directly. Previously you had to write `COPY (SELECT * FROM t) TO ...`, since `COPY t TO` only read the one relation you named. As a bonus, this made logical replication's initial table sync faster too.

- 𝗗 [COPY](https://www.postgresql.org/docs/19/sql-copy.html)
- 𝗖 [7dadd38](https://github.com/postgres/postgres/commit/7dadd38cd), [4c0390a](https://github.com/postgres/postgres/commit/4c0390ac5), [bc2f348](https://github.com/postgres/postgres/commit/bc2f348e8), [2a525cc](https://github.com/postgres/postgres/commit/2a525cc97), [4bea91f](https://github.com/postgres/postgres/commit/4bea91f21)
- 𝗔 Joe Conway, Jian He, Andrew Dunstan, Shinya Kato, Fujii Masao, Kirill Reshke, Ajin Cherian

## Range subtraction with gaps

The `-` operator on ranges has always had an awkward restriction: it errors out if the result would have a hole in the middle, because a single range can't represent two disjoint pieces. The new `range_minus_multi()` returns a *set* of ranges instead:

```sql
-- Subtract a hole out of the middle of a range.
SELECT range_minus_multi('[1,20)'::int4range, '[5,10)'::int4range);
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="plain.sql" output></codapi-snippet>

```text
┌───────────────────┐
│ range_minus_multi │
├───────────────────┤
│ [1,5)             │
│ [10,20)           │
└───────────────────┘
(2 rows)
```

On PostgreSQL 18, `'[1,20)'::int4range - '[5,10)'::int4range` raises `result of range difference would not be contiguous`. There's a multirange version too, which collapses the pieces into one value:

```sql
SELECT multirange_minus_multi(
    '{[1,20)}'::int4multirange,
    '{[5,10)}'::int4multirange
);
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="plain.sql" output></codapi-snippet>

```text
┌────────────────────────┐
│ multirange_minus_multi │
├────────────────────────┤
│ {[1,5),[10,20)}        │
└────────────────────────┘
(1 row)
```

Useful anywhere you compute availability: subtract booked slots from opening hours and get back the gaps.

- 𝗗 [Range Functions](https://www.postgresql.org/docs/19/functions-range.html)
- 𝗖 [5eed8ce](https://github.com/postgres/postgres/commit/5eed8ce50)
- 𝗔 Paul A. Jungwirth

## More jsonpath string methods

`jsonpath` gained eight string methods that mirror their SQL counterparts: `ltrim()`, `rtrim()`, `btrim()`, `lower()`, `upper()`, `initcap()`, `replace()`, and `split_part()`. They chain, so you can clean up values inside the path expression instead of unnesting first:

```sql
SELECT jsonb_path_query_array('["  ada ", "bo smith"]', '$[*].btrim().initcap()');
SELECT jsonb_path_query('"eng-backend-team"', '$.split_part("-", 2)');
SELECT jsonb_path_query('"hello world"', '$.replace("world", "postgres")');
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="plain.sql" output></codapi-snippet>

```text
┌────────────────────────┐
│ jsonb_path_query_array │
├────────────────────────┤
│ ["Ada", "Bo Smith"]    │
└────────────────────────┘
(1 row)

┌──────────────────┐
│ jsonb_path_query │
├──────────────────┤
│ "backend"        │
└──────────────────┘
(1 row)

┌──────────────────┐
│ jsonb_path_query │
├──────────────────┤
│ "hello postgres" │
└──────────────────┘
(1 row)
```

All eight are immutable, like the SQL functions they mirror, which means they're usable in expression indexes.

- 𝗗 [JSON Functions](https://www.postgresql.org/docs/19/functions-json.html)
- 𝗖 [bd4f879](https://github.com/postgres/postgres/commit/bd4f879a9)
- 𝗔 Florents Tselai, David E. Wheeler

## base64url and base32hex

`encode()` and `decode()` learned two more alphabets. `base64url` is the URL-safe variant from [RFC 4648](https://www.rfc-editor.org/rfc/rfc4648.html): `-` and `_` instead of `+` and `/`, no padding, so it's safe in URLs and JWTs without post-processing:

```sql
SELECT encode('\xfbff'::bytea, 'base64')    AS base64,
       encode('\xfbff'::bytea, 'base64url') AS base64url;

SELECT encode('\xdeadbeef'::bytea, 'base32hex') AS base32hex;
SELECT decode('-_8', 'base64url') AS round_trip;
```
<codapi-snippet sandbox="postgres:dev" editor="basic" template="plain.sql" output></codapi-snippet>

```text
┌────────┬───────────┐
│ base64 │ base64url │
├────────┼───────────┤
│ +/8=   │ -_8       │
└────────┴───────────┘
(1 row)

┌───────────┐
│ base32hex │
├───────────┤
│ RQMRTRO=  │
└───────────┘
(1 row)

┌────────────┐
│ round_trip │
├────────────┤
│ \xfbff     │
└────────────┘
(1 row)
```

`base32hex` is the more interesting one: unlike ordinary base32, its alphabet (`0-9`, `A-V`) preserves the sort order of the bytes it encodes, which makes it a compact, sortable encoding for things like UUIDs. One caveat the docs are explicit about: that ordering only holds under a **byte-wise** collation. Sort the encoded text with a natural-language collation and the guarantee is gone, so use `COLLATE "C"` when you rely on it.

- 𝗗 [Binary String Functions](https://www.postgresql.org/docs/19/functions-binarystring.html)
- 𝗖 [497c117](https://github.com/postgres/postgres/commit/497c1170c), [e752a2c](https://github.com/postgres/postgres/commit/e752a2ccc), [e1d9171](https://github.com/postgres/postgres/commit/e1d917182)
- 𝗔 Andrey Borodin, Aleksander Alekseev, Florents Tselai

## Other notable changes

- **CHECK constraints can be un-enforced.** `ALTER TABLE ... ALTER CONSTRAINT ... NOT ENFORCED` now covers CHECK, not just foreign keys, so you can park one while loading awkward data. Re-enforcing re-validates the whole table and fails on any violating row.
- **`GRANT ... GRANTED BY`** records a different role as the grantor, which matters because revoking depends on who granted. `GRANT SELECT ON report TO intern GRANTED BY boss` records `intern=r/boss`, but `boss` must already hold the privilege `WITH GRANT OPTION`.
- **`error_on_null()`** returns its argument or raises, the assertion `coalesce()` can't express: `error_on_null(NULL::int)` fails with `null value not allowed`.
- **`WAIT FOR`** blocks until a standby has replayed to a given LSN: `WAIT FOR LSN '0/1000000' WITH (TIMEOUT '100ms')`. That's read-your-writes on replicas. The default `standby_replay` mode is recovery-only, so on a primary it fails with `recovery is not in progress`.
- **`CHECKPOINT` takes options.** In `CHECKPOINT (MODE SPREAD, FLUSH_UNLOGGED)`, `SPREAD` throttles the checkpoint like a scheduled one instead of flushing as fast as possible (`FAST` is the default).
- **`bytea` ↔ `uuid` casts** work directly now, no more `encode()`/`replace()` gymnastics. Also `tid_block()` and `tid_offset()`, to pull the page number and line pointer out of a `ctid`.
- **Roles and databases as DDL.** `pg_get_role_ddl()`, `pg_get_database_ddl()`, and `pg_get_tablespace_ddl()` reconstruct a definition the way `pg_get_viewdef()` does for views, so recovering one no longer means `pg_dumpall --roles-only`. Every attribute is spelled out, negative ones like `NOSUPERUSER` included; passwords never are.
- **`random(min, max)` for dates and timestamps.** The bounded `random()` from PostgreSQL 17 gained `date`, `timestamp`, and `timestamptz` versions, so test data is a one-liner.
- **`oid8`**, a 64-bit unsigned identifier type. Mostly plumbing for future catalog work, but usable: `'18446744073709551615'::oid8` works where `bigint` overflows.
- **Publications got two new shapes**: `FOR ALL TABLES EXCEPT (TABLE a, TABLE b)` and `FOR ALL SEQUENCES`.
- **Logical replication replicates sequence values**, via `CREATE SUBSCRIPTION`, `REFRESH PUBLICATION`, or the new `REFRESH SEQUENCES`. It also no longer needs `wal_level = logical` up front: `replica` is enough, decoding switches on when something needs it, and the read-only `effective_wal_level` tells you what's in force.
- **`log_min_messages` can be set per process type**: `SET log_min_messages = 'warning, autovacuum:debug1'`. A bare default level is required, otherwise it fails with `Default log level was not defined`.
- **Polish and Esperanto stemmers** for full-text search, plus an updated Dutch one; the old Dutch is still available as `dutch_porter`.
- **Unicode 17.** `unicode_version()` now reports `17.0`, and about 4,800 code points that PostgreSQL 18 considered unassigned are assigned, which matters if you validate with `unicode_assigned()`. `icu_unicode_version()` didn't move, so ICU and the builtin provider can disagree about a brand-new character.
- **List-valued settings can be emptied with `NULL`.** `SET search_path = NULL` now means "empty list", which previously had no spelling at all.

And on the performance and operations side:

- **Foreign key checks got a fast path**, worth up to 2× better insert performance on tables with foreign keys per the [beta 1 announcement](https://www.postgresql.org/about/news/postgresql-19-beta-1-released-3313/).
- **Reads now maintain the visibility map.** On-access pruning sets the all-visible bit in sequential, TID range, sample and bitmap heap scans and the heap side of index scans, not just in `VACUUM` and `COPY ... FREEZE`. On a freshly loaded 256-page table, a plain `SELECT count(*)` marks 221 pages all-visible where PostgreSQL 18 leaves every one unset.
- **`LISTEN`/`NOTIFY` scales to many channels.** A shared channel map means `NOTIFY` wakes only the backends listening on that channel, rather than most of them.
- **More incremental sorts.** `Append` and `MergeAppend` now consider explicit incremental sorts, which mostly helps partitioned tables with an `ORDER BY` that's partially satisfied by an index.
- **Server-side SNI.** A new `pg_hosts.conf`, located by the `hosts_file` setting, maps hostnames to certificate/key pairs, so one server can present a different TLS certificate per requested hostname.
- **New monitoring views.** `pg_stat_lock` reports waits and wait time per lock type; `pg_stat_recovery` exposes replay progress, pause state, and whether promotion was triggered. Many existing views gained a `stats_reset` column, and `pg_stat_progress_vacuum`/`pg_stat_progress_analyze` gained `started_by`, so you can tell an autovacuum from a manual run.
- **`CREATE SUBSCRIPTION ... SERVER`** takes connection details from a `postgres_fdw` foreign server and user mapping instead of an inline connection string, keeping credentials out of the subscription.
- **`postgres_fdw` pushes down more.** Array comparisons in prepared statements now push down, and the new `import_stats` option has `ANALYZE` import remote statistics instead of dragging rows across the wire. Off by default, and spelled `restore_stats` in beta 3.
- **PL/Python can write event triggers**, which previously required PL/pgSQL or C.

## Incompatibilities worth knowing

PostgreSQL 19 changes a handful of defaults and removes some old behavior. The full list is in the [migration section](https://www.postgresql.org/docs/19/release-19.html#RELEASE-19-MIGRATION); these are the ones most likely to surprise you:

- **JIT is off by default.** The cost model deciding when to JIT was unreliable, and a bad decision is expensive. If your analytical queries benefit from JIT, you now have to turn it on deliberately.
- **TOAST compression defaults to `lz4`** when the server is built with `--with-lz4`; otherwise it stays on `pglz`. Existing data isn't recompressed, only newly stored values.
- **`max_locks_per_transaction` doubled to 128** — not a capacity increase. The lock table's shared memory accounting changed (see the hidden gems below), so allocation is deterministic and you hit the limit sooner for a given setting. The rule of thumb from the commit: if you had tuned this, double your value.
- **`log_lock_waits` is on by default**, so waits exceeding `deadlock_timeout` get logged unless you turn it off.
- **`standard_conforming_strings` is forced on** and `escape_string_warning` is gone. Dumps taken by a pre-19 `pg_dump` with it off won't load into a 19 server; re-dump with the newer tool.
- **RADIUS authentication is removed.** It only ever worked over UDP, which can't be secured. There's no drop-in replacement, so plan a migration to LDAP, OAuth, or certificates.
- **MD5 password authentication warns on every success.** `md5_password_warnings` isn't new (18 added it, to warn when a password was *set*); 19 reuses it for logins, the next nudge toward [`scram-sha-256`](https://www.postgresql.org/docs/19/auth-password.html), the challenge-response scheme ([RFC 7677](https://datatracker.ietf.org/doc/html/rfc7677)) that replaces it. Related: `password_expiration_warning_threshold` (default 7 days) warns before a password expires.
- **`json_array()` over an empty subquery returns `[]`**, not NULL.
- **The built-in GiST `inet_ops` is now the default opclass for `inet`/`cidr`**, displacing `btree_gist`'s `gist_inet_ops` and `gist_cidr_ops`, which could miss rows they should have returned. `pg_upgrade` refuses a cluster with an index on the old opclasses, and reindexing won't help: drop and recreate them on `inet_ops`. It likewise refuses `MULE_INTERNAL` encoding (removed) and carriage returns or line feeds in database, role, or tablespace names (now disallowed).
- **`pg_stat_subscription_stats.sync_error_count` is renamed** to `sync_table_error_count`, since sequence sync errors are counted separately now. The `BUFFERPIN` wait event type is renamed to `BUFFER`.

## Tooling

A grab bag of client and utility improvements:

- **`psql` can display booleans however you like**: `\pset display_true yes` and `\pset display_false no`, instead of the terse `t`/`f`.
- **Two new prompt escapes.** `%S` shows the current `search_path` (it needs an 18-or-later server), and `%i` shows whether you're on a hot standby — a nice guard against writing to the wrong host.
- **`\dRp+`, `\dRs+`, and `\dX+` show comments** for publications, subscriptions, and extended statistics, plus a large batch of tab-completion improvements, including one for `FOR PORTION OF`.
- **`vacuumdb --dry-run`** prints the commands it would run instead of running them. Also, `--analyze-only` and `--analyze-in-stages` no longer skip partitioned tables.
- **`pgbench --continue-on-error`** keeps a run going after SQL errors instead of aborting the client, which makes benchmarking workloads with expected conflicts far less annoying.
- **`pg_test_timing` reports nanoseconds** instead of microseconds, and adds a table of exact timings alongside the histogram (with an optional `--cutoff`). The new `timing_clock_source` setting can select the TSC directly on x86, making `EXPLAIN (ANALYZE, TIMING)` cheaper.
- **`pg_waldump` and `pg_verifybackup` read WAL from tar archives**, so you no longer have to extract a backup to inspect it.
- **`pg_upgrade` is much faster with many large objects**, the pathological case that made some upgrades take hours. It also handles non-default tablespaces stored inside `PGDATA` instead of erroring out.
- **`pg_dump` can dump restorable extended statistics**, so a restored database doesn't start out with none for its `CREATE STATISTICS` objects.
- **"Grease mode" ran during the beta, and then stops.** libpq gained protocol parameters that deliberately exercise unknown-parameter handling, to smoke out proxies and drivers that mishandle protocol extensions. It was active during the 19 beta and will **not** be sent as of the release, though the parameter reservations stay for future use. If you maintain a connection pooler or a wire-protocol driver, this is the release to test against.

## Hidden gems

Everything above comes from the release notes. Those notes are deliberately curated: around 3,400 commits landed between PostgreSQL 18.0 and the 19 stable branch, and a release note that listed all of them would be useless. But a few of the changes that didn't make the cut are still interesting, so here are some worth knowing about:

- **You can finally find out who killed your backend.** When a session is terminated by `pg_terminate_backend()` or an external `SIGTERM`, the server log now adds `DETAIL: Signal sent by PID 142, UID 999.` It's an `errdetail_log()`, so it goes to the log rather than the client, and it needs `SA_SIGINFO` (Linux, FreeBSD, most modern Unixes). A follow-up reworked signal handling to pass a `pg_signal_info` struct into every handler, instead of stashing the sender in globals.
𝗖 [55890a9](https://github.com/postgres/postgres/commit/55890a91945), [3e2a149](https://github.com/postgres/postgres/commit/3e2a1496bae) • 𝗔 Jakub Wartak, Andrew Dunstan

- **A quiet rewrite of shared memory allocation.** The release notes record the user-visible part, the new `ShmemRequestStruct()` API. Behind it, Heikki Linnakangas converted essentially every subsystem: the buffer manager, AIO, SLRUs, `lwlock.c`, `predicate.c`, `pg_stat_statements`. The fudge factors came out too — a 10% "safety margin" in the lock manager's hash table estimates, a bogus one in `predicate.c` — which is exactly why `max_locks_per_transaction`'s default had to double. The commit spells it out: the allocation became deterministic, "but it also means that you often hit one of the limits sooner than before".
𝗖 [283e823](https://github.com/postgres/postgres/commit/283e823f9), [a4b6139](https://github.com/postgres/postgres/commit/a4b6139dcce), [9b5acad](https://github.com/postgres/postgres/commit/9b5acad3f40), [3e854d2](https://github.com/postgres/postgres/commit/3e854d2ff1a), [79534f9](https://github.com/postgres/postgres/commit/79534f906)

- **EXPLAIN tells you what a `Result` node replaced.** A plan proven empty at planning time used to collapse to a bare `Result` with `One-Time Filter: false`, losing all trace of the relation it stood in for. Now it says `Replaces: Scan on w`, which makes "why is my table missing from this plan" answerable.
𝗖 [f2bae51](https://github.com/postgres/postgres/commit/f2bae51dfd5)

- **CI moved off Cirrus CI onto GitHub Actions.** Not a database change, but it's the infrastructure every future patch is tested on. Cirrus support was removed outright rather than kept in parallel.
𝗖 [9c12606](https://github.com/postgres/postgres/commit/9c126063b19), [68c8a36](https://github.com/postgres/postgres/commit/68c8a365d4d) • 𝗔 Andres Freund

- **Set-returning functions are no longer allowed in a window `OVER` clause.** An SRF there contradicts the principle that a window function doesn't change the row count, and drew two bug reports (#17502 and #19535). Rather than define the semantics, 19 makes it an error; put the SRF in a `LATERAL` `FROM` clause instead. Worth flagging because it's a hard error on code that previously "worked".
𝗖 [0c15b71](https://github.com/postgres/postgres/commit/0c15b715c65) • 𝗔 Tom Lane

- **The multixact members offset is 64-bit.** Widening `MultiXactOffset` lifts the 2^32 cap on total multixact members and removes members-space wraparound, along with the emergency anti-wraparound freezing that came from exhausting it; multixact IDs themselves are still capped at 2^31. The on-disk format changed, so `pg_upgrade` rewrites the `pg_multixact` files. Related: the wraparound warning threshold moved from 40 million to 100 million transactions.
𝗖 [bd8d9c9](https://github.com/postgres/postgres/commit/bd8d9c9bd), [48f11bf](https://github.com/postgres/postgres/commit/48f11bfa0) • 𝗔 Maxim Orlov, Nathan Bossart

- **SIMD keeps spreading.** `COPY FROM` parsing for text and CSV, `hex_encode()`/`hex_decode()`, page checksums (AVX2), and CRC32C on ARM all moved to vector instructions. No API changes; bulk loading and checksumming just get faster.
𝗖 [e0a3a3f](https://github.com/postgres/postgres/commit/e0a3a3fd5), [ec8719c](https://github.com/postgres/postgres/commit/ec8719ccb), [5e13b0f](https://github.com/postgres/postgres/commit/5e13b0f24), [fbc57f2](https://github.com/postgres/postgres/commit/fbc57f2bc)

- **AIX support is back**, after being dropped in PostgreSQL 17. IBM stepped up with buildfarm animals; xlc and 32-bit builds did not come back, so it's gcc-only and 64-bit-only. Meanwhile the minimum C version moved from C99 to C11, Visual Studio 2019 is now the floor on Windows, and MSVC can build for AArch64.
𝗖 [4a1b05c](https://github.com/postgres/postgres/commit/4a1b05caa), [ecae097](https://github.com/postgres/postgres/commit/ecae09725), [f5e0186](https://github.com/postgres/postgres/commit/f5e0186f8), [8fd9bb1](https://github.com/postgres/postgres/commit/8fd9bb1d9), [a516b3f](https://github.com/postgres/postgres/commit/a516b3f00)

## Final thoughts

PostgreSQL 19 is a big release, and unusually front-loaded with things you can *see*. A few themes stand out:

- **Query language:** `FOR PORTION OF`, `ON CONFLICT DO SELECT`, and `IGNORE NULLS` are small additions, but each removes a well-known workaround.
- **Operations:** `REPACK CONCURRENTLY` and autovacuum prioritization both attack the same problem, which is that maintenance used to mean downtime.
- **Planner:** eager aggregation and `NOT IN` anti-joins are real wins on real queries, and `pg_plan_advice` is a notable philosophical shift for a project that spent twenty years saying no to hints.
- **Observability:** `EXPLAIN (ANALYZE, IO)` finally makes the asynchronous I/O added in 18 visible, and Memoize estimates let you check the planner's reasoning rather than just its conclusion.

The defaults changed more than usual too. JIT off, `lz4` TOAST compression, `log_lock_waits` on, `max_locks_per_transaction` doubled, RADIUS gone: read the [migration notes](https://www.postgresql.org/docs/19/release-19.html#RELEASE-19-MIGRATION) before you upgrade, not after.

P.S. Curious how we handle time series data at scale? [VictoriaMetrics](https://github.com/VictoriaMetrics) is a purpose-built database for metrics, logs, and traces. Browse the [rest of our blog](https://victoriametrics.com/blog) for deep dives into storage engines, query performance, and observability at scale.
