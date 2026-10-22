---
draft: false
page: blog blog_post
authors:
  - Michelle Sebek
date: 2026-09-07
title: "Two cats, two dogs, four vendors, and the model the AI couldn't find (Tech Talk Companion)"
summary: "Tech Talk #13: Mathias runs his house and four pets on Home Assistant, VictoriaMetrics, and vmanomaly. Then the AI picked the wrong model. Watch on demand."
enableComments: true
categories:
  - Tech Talk
tags:
  - VictoriaMetrics
  - Tech Talk
  - vmalert
  - vmanomaly
  - Home Assistant
  - Video
images:
  - /blog/tech-talk-streems/observability-pet-project/ai-skills-victoriametrics-replay-v2.webp
---

Tech Talks went dark for a few months, and on [episode 13](https://www.youtube.com/watch?v=n_aN4Z2Eb8I) I finally got to ask why. Mathias Palmersheim's answer, delivered completely straight, was that his users were unhappy with the availability and usability of their feeders and their litter box, and he wasn't allowed back on stream until that got fixed.

The users are two dogs and two cats, and they have titles. Maisie, a Shiba Inu who came to him through a rescue, is the recently promoted chief executive pawofficer. Millie is head of `/dev/null` as a service, on the grounds that she's a void that consumes everything and somehow never gains an ounce. Max is doing his best. Moxxie is the chaos engineer, and she's the one who matters later.

This was my first tech talk since joining VictoriaMetrics. I came in expecting a cute demo and got a better argument about alert design than I've heard from most people talking about production systems.

## The solutions engineering team lives in Home Assistant

Mathias's counterpart Zakhar has a coffee maker wired into [Home Assistant](https://www.home-assistant.io/), and Home Assistant knows when he's actually awake. Most coffee makers let you set an alarm, so 7:45 on the dot, every day, including Saturday, and by the time you get downstairs the burner has been cooking it for an hour. Zakhar's coffee maker brews when he's up and moving instead. Stephen, another teammate, went further and built himself an Iron Man wake-up routine, where the lights come up gradually as a gentle alarm and he gets a report on the day when he opens his eyes. Nobody on that team is getting jolted awake by the same song every morning.

I mention this because it explains the rest of the talk. These are people who use this stuff at home for reasons that have nothing to do with work, which is why the pet project isn't really a toy.

## Many vendors, one pane of glass, and a hard rule about the cloud

The device shopping came with a constraint that'll be familiar to anyone who's worked a support rotation through a cloud outage. Mathias wanted his lights to come on when us-east-1 is having a bad day, and he was blunt about why: that's exactly when he's busiest helping customers whose own systems are down. So every device had to be controllable over the LAN, which knocks out a lot of otherwise nice hardware.

No single vendor makes sensors for everything, either. PetLibro does the fountains and feeders, Whisker makes the Litter-Robot, Roborock cleans up afterward, and TP-Link and Aqara handle the lights and switches. Four vendors, four apps, four data models, and a homeowner who does not want to open four apps to find out if the kids left the lights on or if someone forgot to take care of the animals.

Home Assistant is the glue that holds all of this together. Sensors push to it, or it polls a cloud API for the few devices that only speak cloud. Then comes the part I didn't expect. Mathias called Home Assistant his favorite Prometheus exporter, and he meant that literally: the [Prometheus integration](https://www.home-assistant.io/integrations/prometheus/) publishes everything Home Assistant knows at `/api/prometheus`, behind a long-lived access token.

After that it's an ordinary scrape. He uses Telegraf rather than a scrape job, which needs the format declared explicitly because Telegraf speaks several, plus a field drop filter doing the same job as a relabeling rule. Home Assistant emits state-change counters and metadata he has no use for, and that metadata is high cardinality, so he drops it at the source rather than storing it and filtering later. Telegraf pushes into a single-node VictoriaMetrics, and Grafana reads from there.

Why bother, when Home Assistant already has dashboards? Because long ranges are awkward in them, and comparing a few metrics over the same window is worse. In Grafana he pulled up 90 days of cat weight, and Moxxie's line is drifting up. She seems to be getting ready for winter despite it being 100 degrees outside where she lives.

## The alert that clears itself

The second reason is a design argument, and it's the part of the talk I'd repeat to people building alerting for a team.

Home Assistant has a powerful automation engine, and it is amazing for things like turning on the lights when a motion sensor is tripped. But for being notified of an issue that requires human intervention—like filling an empty water bowl or dealing with a full litter box—Mathias wanted something similar to an incident response tool that updates or deletes notifications once they are resolved, so he didn't have to check a dashboard to see if someone else had already resolved the issue. The other thing that Home Assistant notifications were lacking was the `ALERTS` metric, which is how vmalert tracks which rules are in a pending or firing state over time. This metric is the backbone of the [alert statistics dashboard](https://grafana.com/grafana/dashboards/24553-victoriametrics-alert-statistics/) maintained by VictoriaMetrics, which Mathias uses to show his stakeholders/pets that they are being well taken care of so they will allow him to continue doing tech talks.

So the metrics live in VictoriaMetrics, [vmalert](https://docs.victoriametrics.com/victoriametrics/vmalert/) evaluates alerting rules and sends alerts to Alertmanager, Alertmanager posts a webhook to a self-hosted [ntfy](https://docs.ntfy.sh/) server, and everyone in the house gets one push. He admitted to using AI to write a small service for deleting notifications from ntfy after Alertmanager resolves them. If you'd rather not write that glue, Alertmanager also supports email, Discord, Telegram, and more or less every incident response tool there is.

The alert we watched was called `Thicc Kitty`. Mathias had quietly moved the threshold down to 10.5 pounds so it would trigger during the demo, and I caught him, because the dashboard right next to it said 11. He copped to it. The notification lands on his phone with a link straight back to the Grafana panel, so at least the alert about his cat's weight is actionable.

## Thresholds in a git repo, edited by someone who doesn't use an IDE

This is the trick I'd steal outright, and it's one flag.

vmalert loads rules from plain files, Kubernetes resources, or an HTTP URL. Mathias points it at a git repo over HTTP. That means the threshold for `Thicc Kitty` is one line in a file with a web editor sitting in front of it, and his wife doesn't have to live in an IDE to change it. She opens the repo in a browser, changes 10.5 to 11.25, commits it, and vmalert picks it up on its next poll—once a minute in his setup—and we sat through that minute on stream twice waiting for the alert to resolve. There was a request for dramatic music.

Grafana shows these as [data source managed alerts](https://docs.victoriametrics.com/guides/vmalert-datasource-managed-alerts-grafana/) by querying the vmalert API through the data source config. If you'd rather skip Grafana, setting `-vmalert.proxyURL` also surfaces the same rules and states in VMUI. Stephen described the route he prefers in chat. He applies the operator's [VMRule](https://docs.victoriametrics.com/operator/resources/vmrule/) resource to the cluster, which writes the rules out as a plain file with the same hot reload, and Argo CD pulls every minute. Browser tab or GitOps, same idea.

And the idea is worth saying plainly, because it generalizes past cats. The threshold is the piece of an alerting rule most likely to be wrong and least likely to need an engineer to fix. Put it somewhere the person who actually knows the right number can reach it. In this house that person is his wife, relaying what the vet said. At work it's probably your application/service owner, and either way you've taken yourself out of the loop for the change that happens most often.

## Where static thresholds fall short

The vet gave him a number for weight. Above 11.25 pounds, Moxxie needs more play, less food, or different food. That's the easy case, and it's rarer than you'd think.

The vet did not give him a number for litter box visits. Too few can mean the cat is backed up or can't get to the box. Too many can point at an infection. Both directions are bad, neither has a published figure, and the normal range depends on the individual animal. Mathias's summary was that cats are too chaotic for static thresholds most of the time, which is how Moxxie ended up with the chaos engineer title.

He built the query as a 24-hour max, averaged to smooth it, one series per cat. Just eyeballing the graph, Millie had a couple of multi-day stretches with no visits at all, and Moxxie had a spike. Catch both of those with a rule that doesn't also catch normal variation and you've written the kind of interconnected rule set nobody on the team will touch a year from now.

[vmanomaly](https://docs.victoriametrics.com/anomaly-detection/) is the answer here, and Fred Navruzov, who leads it, sat down with Mathias before the stream to tune the model. They used [Temporal Envelope](https://docs.victoriametrics.com/anomaly-detection/components/models/#temporal-envelope), an in-house model added in v1.30.0 that learns an expected range from trend, calendar patterns and holidays, and keeps adapting during inference without letting one spike bend it permanently. Z-score and median absolute deviation don't handle seasonality at all. Prophet handles seasonality but does not support online updates to the model; Temporal Envelope supports both seasonality and online model updates.

Tuning is where your own knowledge goes in, and the parameters stay readable even when the statistics behind them don't. The ones they reached for:

- [`detection_direction`](https://docs.victoriametrics.com/anomaly-detection/components/models/#detection-direction), set to both here, since either extreme means something. For request latency you'd set it to above only. Mathias put it better: nobody has ever been upset that a website loaded too fast.
- [`clip_predictions`](https://docs.victoriametrics.com/anomaly-detection/components/models/#clip-predictions) pins the model to a plausible range. The floor is 0.001, because you can't use a litter box a negative number of times, but using the litter box 0 times for 2 days should still be flagged as an anomaly.
- [`min_deviation_from_expected`](https://docs.victoriametrics.com/anomaly-detection/components/models/#minimal-deviation-from-expected) throws out deviations too small to care about.
- Seasonality presets handle hour-of-day and day-of-week shapes, and you can encode holidays directly, which is how you stop Black Friday from reading as an incident.
- A streaks setting makes a condition persist before it counts, so one lonely spike doesn't wake anybody up.

The UI plots the model against your real data while you turn these knobs, so you watch the expected range move. Once it looked right they exported it, either as a full standalone config with server info and queries or as just the model block. Mathias took the model block, since he already had vmanomaly running.

## The uncomfortable part

Fred and Mathias needed about an hour. I asked how long the whole build took, and he raised the harder question himself: what do you do if you don't have a Fred (our product lead)?

vmanomaly ships [an assistant in the UI](https://docs.victoriametrics.com/anomaly-detection/ui/#ai-assistance) for exactly that. It takes any OpenAI-compatible endpoint, Anthropic-compatible endpoints, OpenRouter, or Amazon Bedrock, and the model behind those endpoints can drive the vmanomaly UI. In this case Mathias and Fred were running Claude Fable 5. He typed a plain-English prompt asking for a model that uses streaks to spot litter box use that's too frequent or not frequent enough, approved the tool calls, and we watched it drive the interface with results streaming back.

It picked the Prophet model for everything, and Mathias recognized the failure on sight because he and Fred had already hit it while building the [agent skills](https://docs.victoriametrics.com/ai-tools/#agent-skills). The context the model was reading from pushed it toward the Prophet model no matter what the data looked like, and he'd forgotten to update his skills before going live. Mathias had also forgotten to add the MCP server URL for the vmanomaly assistant, so the assistant was working with one hand tied behind its back. This misconfiguration led the vmanomaly assistant to believe it didn't have access to the new `temporal_envelope` model, so the assistant settled for the seasonal quantile model instead. The seasonal quantile model landed reasonably close, caught the spike and the trough, and never found Temporal Envelope at all—an issue that might be addressed in next month's tech talk.

Here is what I did not expect on a vendor livestream. Mathias said it out loud: our expertise is better than what you can get with just Claude, and we confirmed that today.

Even after Mathias had fixed the MCP issue and updated his skills after the stream, the Temporal Envelope model that Fable came up with was still a bit worse than what Mathias and Fred had built without the assistant. Chat sharpened it further. Mark pointed out you should define the agent's role and its guardrails before you ask it anything, which is the fix for a model confidently grabbing the wrong tool. Stephen pointed out that a UI showing you what the model is doing makes the model usable by almost anyone. He's right, and it's why the wrong answer was obvious inside of a minute instead of after a week of bad alerts. An assistant that hands you a config you can't inspect is a much riskier proposition than one whose output you can see drawn over your own data.

## What it costs to run

Metrics, logs, traces, [vmanomaly](https://victoriametrics.com/products/enterprise/anomaly-detection/), and all the glue between them run on two CPUs and four to six gigabytes of RAM in his lab. Machine learning without a rack of H100s in the basement, as he put it. If you do have a rack of servers in the basement—or, more likely, a data center that needs watching—the same stack scales with the Kubernetes operator and the cluster version.

Most of what's above is the open source core. vmanomaly is part of our enterprise offering and needs a license key, and it reads from and writes to whichever VictoriaMetrics you're already running, open source or enterprise, with vmalert and Grafana untouched. There's a [trial license](https://victoriametrics.com/products/enterprise/trial/) if you want to try the litter box thing at home.

The whole thing, tuned and shipped, took him under 90 minutes.

## Now I want one for Gizmo

My cat is an outdoor cat who lives in our garage. She sleeps there, eats there, and conducts her other business there, and I have no idea what she does all day. I want a GoPro on it and an alert when there's drama.

There's a real engineering problem in the way, though, and it's the kind this talk was good at surfacing. Gizmo is heavier in winter and lighter in summer when she sheds, so her weight swings for reasons that have everything to do with survival and nothing to do with her health. I have no way to put her on a scale in the first place. Mathias thinks that's a follow-up stream. Fred, start warming up.

Next month Roman, our co-founder, takes Tech Talk #14 through VictoriaMetrics, VictoriaLogs and VictoriaTraces with an AI agent, using the skills in our MCP server across the whole stack instead of one component's UI. Mark's question about guardrails gets answered properly, on stream, by someone who wrote the skills.

![Gizmo, an outdoor cat who resides in the garage](/blog/tech-talk-streems/observability-pet-project/gizmo.webp)

<figcaption style="text-align: center; font-style: italic;">Gizmo, our outdoor cat who resides in the garage</figcaption>

We'll have a booth at WeAreDevelopers at the end of September, All Things Open in Raleigh in October, and KubeCon in November. There's a [Google Form](https://docs.google.com/forms/d/e/1FAIpQLSeshaBBdN3weXMcSYD7_xfvIssx0rj_0t00AWsXf34Uw2Mysw/viewform?usp=header) in the video description for suggesting topics, and I read everything that comes to [michelle@victoriametrics.com](mailto:michelle@victoriametrics.com). If you've found a use for VictoriaMetrics we haven't thought of, send it.

Also, settle something for me: how many pets, and are you a dog person or a cat person? Chat was split.

*By engineers, for engineers.*
