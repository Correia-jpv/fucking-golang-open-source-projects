# Removal and Migration Log

[Chinese version](移除与迁移记录.md)

This document records the most important repository migrations and removals made during this refactor so future maintainers can trace the decisions.

## Migrated to New Repository Entries

| Old entry | New entry | Notes |
| --- | --- | --- |
| `docker/docker` | `moby/moby` | The Docker engine upstream source now lives under Moby. |
| `vmware/harbor` | `goharbor/harbor` | Harbor's official repository has moved. |
| `coreos/clair` | `quay/clair` | Clair has moved to the Quay organization. |
| `containous/traefik` | `traefik/traefik` | Traefik's official repository name was updated. |
| `lonelycode/tyk` | `TykTechnologies/tyk` | Tyk moved to its official organization. |
| `bitly/nsq` | `nsqio/nsq` | NSQ moved to its official repository. |
| `youtube/vitess` | `vitessio/vitess` | Vitess changed its official organization. |
| `ipfs/go-ipfs` | `ipfs/kubo` | Go IPFS was renamed to Kubo. |
| `henrylee2cn/pholcus` | `andeya/pholcus` | Pholcus now uses this official repository entry. |
| `astaxie/beego` | `beego/beego` | Current maintenance entry for Beego. |
| `robfig/revel` | `revel/revel` | Current maintenance entry for Revel. |
| `micro/micro` | `micro/go-micro` | Current source entry for the microservices framework. |
| `drone/drone` | `harness/harness` | Drone has become part of the Harness ecosystem. |
| `etsy/hound` | `hound-search/hound` | Hound moved to its official repository. |
| `mattermost/platform` | `mattermost/mattermost` | Current main repository for Mattermost. |

## Archived and Removed Directly

The following projects were clearly archived at review time, so they are no longer included:

- `coreos/rkt`
- `shipyard/shipyard`
- `zettio/weave`
- `coreos/torus`
- `sourcegraph/appdash`
- `uber/tchannel`
- `bosun-monitor/bosun`
- `gravitational/satellite`
- `ooyala/atlantis`
- `youtube/doorman`
- `nanopack/yoke`
- `silenceper/dcmp`
- `inconshreveable/ngrok`
- `yahoo/gryffin`
- `apex/apex`
- `codeskyblue/gosuv`
- `vzex/dog-tunnel`
- `pinggg/pingd`
- `gernest/utron`
- `lunny/tango`

## Long-Inactive or Overlapping Projects No Longer Recommended

The following projects did not enter the refreshed catalog mainly because they are long-inactive, overlap heavily with more active projects, or are now difficult to recommend as first-choice options:

- `huichen/wukong`
- `codetainerapp/codetainer`
- `andyxning/shortme`
- `afex/hystrix-go`
- `mehrdadrad/mylg`
- `cyfdecyf/cow`
- `huichen/mlf`
- `Qihoo360/poseidon`
- `idcos/osinstall`
- `fagongzi/gateway`
- `andot/hprose`
- `Terry-Mao/bfs`
- `TalkingData/owl`
- `cloudinsight/cloudinsight-agent`
- `rach/pome`
- `open-falcon/of-release`
- `mesos/cloudfoundry-mesos`
- `laincloud/lain`
- `weibocom/opendcp`
- `hoisie/web`
- `codegangsta/martini`
- `goshawkdb/server`
- `slicebit/qb`
- `chasex/redis-go-cluster`
- `siddontang/ledisdb`
- `degdb/degdb`
- `blackbeans/kiteq`

## Notes

- This refresh is not a simple star-count ranking. The catalog was rebuilt around whether a project is still worth recommending today.
- A small number of still-usable but clearly slower-moving projects are not all listed here; they have naturally exited the main catalog in the refreshed version.
