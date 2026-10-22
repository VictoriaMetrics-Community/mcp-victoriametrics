---
build:
  list: never
  publishResources: false
  render: never
sitemap:
  disable: true
---

* Cut release in changelog.

* Create tag locally:
```sh
git tag v0.1.8 -m "Release v0.1.8" -s 
```

* Publish release:
```sh
TAG=v0.1.8 make publish-release
```

* Create draft Github Release:
```sh
TAG=v0.1.8 make github-create-release
```

* Upload Release assets:
```sh
TAG=v0.1.8 make github-upload-assets
```

* Publish [release on Github](https://github.com/VictoriaMetrics/vmestimator/releases).