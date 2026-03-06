# Golang Open Source Projects

面向中文读者重新整理的 Go 开源项目目录。新版目录不再追求“尽可能全”，而是优先保留仍在维护、社区认知清晰、适合学习和选型的项目，并补充了 AI Agent 相关项目。

当前版本收录 **75** 个项目，分成 **10** 个主题；最近一次维护状态审阅时间为 **2026-03-06**。

- [分类与维护策略](docs/分类与维护策略.md)
- [移除与迁移记录](docs/移除与迁移记录.md)

## 这次整理做了什么

- 把旧的 17 个松散分类重组为 10 个主题，去掉了难维护的“其它”分类。
- 清理了已归档、仓库已废弃、长期停更且已有明确替代方案的项目。
- 去掉了重复收录，同一个项目只保留一个最合适的入口分类。
- 新增 AI / Agent 分类，覆盖 LLM 应用框架、MCP、推理运行时和向量检索。

## 收录原则

- 优先保留截至 2026-03-06 仍可确认处于维护状态的项目。
- 已归档、仓库消失或长期停更且已有更好替代的项目默认移除。
- 目录强调学习与工程选型价值，不再为了覆盖面保留大量边缘项目。
- 每个项目只收录一次，避免在多个分类中重复出现。

## 分类导航

| 分类 | 关注点 | 项目数 |
| --- | --- | --- |
| AI / Agent | LLM 应用框架、MCP、模型运行时与向量能力 | 7 |
| 云原生与容器 | 容器运行时、编排、镜像仓库和集群平台 | 8 |
| 服务治理与平台工程 | PaaS、服务治理、CI/CD、消息与异步任务 | 12 |
| 数据存储与搜索 | 数据库、分布式存储、检索与数据访问生态 | 10 |
| 可观测性 | 指标、图表、告警与运行状态检查 | 5 |
| 网络与安全 | 网关、负载均衡、代理、流量调试与网络工具 | 6 |
| Web 开发与应用 | Web 框架、服务端组件与实时交互能力 | 11 |
| 数据处理与机器学习 | ML、NLP、爬虫与数据处理 | 6 |
| 开发者工具与基础库 | 开发效率、测试、终端 UI 和核心基础库 | 8 |
| 区块链 | 仍在维护、影响力最大的 Go 区块链项目 | 2 |

## AI / Agent

LLM 应用框架、MCP、模型运行时与向量能力

| 项目 | 简介 |
| --- | --- |
| <b><code>164240⭐</code></b> <b><code>&nbsp;14793🍴</code></b> [ollama/ollama](https://github.com/ollama/ollama)) | 本地运行、分发和管理大模型的 Go 运行时。 |
| <b><code>&nbsp;&nbsp;8775⭐</code></b> <b><code>&nbsp;&nbsp;1045🍴</code></b> [tmc/langchaingo](https://github.com/tmc/langchaingo)) | Go 版 LLM 应用框架，覆盖 prompt、tool calling、agent 和 RAG。 |
| <b><code>&nbsp;&nbsp;9884⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;763🍴</code></b> [cloudwego/eino](https://github.com/cloudwego/eino)) | CloudWeGo 出品的 Go AI 应用框架，强调组件化编排和生产落地。 |
| <b><code>&nbsp;&nbsp;8285⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;781🍴</code></b> [mark3labs/mcp-go](https://github.com/mark3labs/mcp-go)) | 用 Go 构建 MCP client 和 server 的实用 SDK。 |
| <b><code>&nbsp;43312⭐</code></b> <b><code>&nbsp;&nbsp;3640🍴</code></b> [mudler/LocalAI](https://github.com/mudler/LocalAI)) | OpenAI 兼容的本地推理服务，适合私有化部署。 |
| <b><code>&nbsp;&nbsp;1634⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;245🍴</code></b> [mudler/LocalAGI](https://github.com/mudler/LocalAGI)) | 面向本地模型的 Agent 平台，强调工具调用和自治流程。 |
| <b><code>&nbsp;15741⭐</code></b> <b><code>&nbsp;&nbsp;1203🍴</code></b> [weaviate/weaviate](https://github.com/weaviate/weaviate)) | Go 编写的向量数据库，可用于 RAG、检索和 Agent memory。 |

## 云原生与容器

容器运行时、编排、镜像仓库和集群平台

| 项目 | 简介 |
| --- | --- |
| <b><code>&nbsp;71484⭐</code></b> <b><code>&nbsp;18913🍴</code></b> [moby/moby](https://github.com/moby/moby)) | Docker 引擎的上游项目，也是学习容器运行时实现的核心入口。 |
| <b><code>120953⭐</code></b> <b><code>&nbsp;42598🍴</code></b> [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes)) | 事实标准级的容器编排平台。 |
| <b><code>&nbsp;27675⭐</code></b> <b><code>&nbsp;&nbsp;5123🍴</code></b> [goharbor/harbor](https://github.com/goharbor/harbor)) | 企业级 OCI 镜像仓库，带权限、审计和复制能力。 |
| <b><code>&nbsp;25398⭐</code></b> <b><code>&nbsp;&nbsp;3172🍴</code></b> [rancher/rancher](https://github.com/rancher/rancher)) | 面向多集群场景的 Kubernetes 管理平台。 |
| <b><code>&nbsp;10940⭐</code></b> <b><code>&nbsp;&nbsp;1195🍴</code></b> [quay/clair](https://github.com/quay/clair)) | 容器镜像漏洞分析与扫描服务。 |
| <b><code>&nbsp;&nbsp;3588⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;658🍴</code></b> [moby/swarmkit](https://github.com/moby/swarmkit)) | Docker Swarm 的核心编排组件，适合学习调度和集群编排。 |
| <b><code>&nbsp;&nbsp;4654⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;941🍴</code></b> [AliyunContainerService/pouch](https://github.com/AliyunContainerService/pouch)) | 阿里开源的容器引擎项目，聚焦更强的隔离与稳定性。 |
| <b><code>&nbsp;16255⭐</code></b> <b><code>&nbsp;&nbsp;2056🍴</code></b> [hashicorp/nomad](https://github.com/hashicorp/nomad)) | 轻量级工作负载编排器，适合对比 Kubernetes 的另一条路线。 |

## 服务治理与平台工程

PaaS、服务治理、CI/CD、消息与异步任务

| 项目 | 简介 |
| --- | --- |
| <b><code>&nbsp;&nbsp;5257⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;550🍴</code></b> [tsuru/tsuru](https://github.com/tsuru/tsuru)) | 成熟的开源 PaaS，适合学习应用平台抽象。 |
| <b><code>&nbsp;&nbsp;6116⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;855🍴</code></b> [goodrain/rainbond](https://github.com/goodrain/rainbond)) | 以应用为中心的云原生平台，覆盖交付、运维和微服务治理。 |
| <b><code>&nbsp;33854⭐</code></b> <b><code>&nbsp;&nbsp;2896🍴</code></b> [harness/harness](https://github.com/harness/harness)) | Drone 已并入 Harness 生态后，新的 CI/CD 与开发者平台入口。 |
| <b><code>&nbsp;19957⭐</code></b> <b><code>&nbsp;&nbsp;2004🍴</code></b> [gravitational/teleport](https://github.com/gravitational/teleport)) | 基于零信任模型的远程访问与基础设施入口。 |
| <b><code>&nbsp;38074⭐</code></b> <b><code>&nbsp;&nbsp;8258🍴</code></b> [istio/istio](https://github.com/istio/istio)) | 服务网格代表项目，覆盖流量治理、安全和可观测性。 |
| <b><code>&nbsp;&nbsp;&nbsp;&nbsp;10⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [uber/jaeger](https://github.com/uber/jaeger)) | 分布式追踪系统，适合与 OpenTelemetry 一起理解链路追踪。 |
| <b><code>&nbsp;27581⭐</code></b> <b><code>&nbsp;&nbsp;2451🍴</code></b> [go-kit/kit](https://github.com/go-kit/kit)) | Go 微服务开发工具箱，强调可观测性和可测试性。 |
| <b><code>&nbsp;&nbsp;6058⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;580🍴</code></b> [goadesign/goa](https://github.com/goadesign/goa)) | 设计优先的 Go 服务开发框架。 |
| <b><code>&nbsp;10658⭐</code></b> <b><code>&nbsp;&nbsp;1160🍴</code></b> [TykTechnologies/tyk](https://github.com/TykTechnologies/tyk)) | 成熟的开源 API Gateway。 |
| <b><code>&nbsp;22722⭐</code></b> <b><code>&nbsp;&nbsp;2403🍴</code></b> [micro/go-micro](https://github.com/micro/go-micro)) | Go 微服务框架，适合研究服务抽象与插件化扩展。 |
| <b><code>&nbsp;25881⭐</code></b> <b><code>&nbsp;&nbsp;2908🍴</code></b> [nsqio/nsq](https://github.com/nsqio/nsq)) | 经典的实时分布式消息平台。 |
| <b><code>&nbsp;&nbsp;7950⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;944🍴</code></b> [RichardKnop/machinery](https://github.com/RichardKnop/machinery)) | Go 异步任务队列，适合替代 Celery 的思路参考。 |

## 数据存储与搜索

数据库、分布式存储、检索与数据访问生态

| 项目 | 简介 |
| --- | --- |
| <b><code>&nbsp;31995⭐</code></b> <b><code>&nbsp;&nbsp;4087🍴</code></b> [cockroachdb/cockroach](https://github.com/cockroachdb/cockroach)) | 分布式 SQL 数据库，强调强一致与弹性扩展。 |
| <b><code>&nbsp;20846⭐</code></b> <b><code>&nbsp;&nbsp;2311🍴</code></b> [vitessio/vitess](https://github.com/vitessio/vitess)) | YouTube 开源的 MySQL 水平扩展方案。 |
| <b><code>&nbsp;39870⭐</code></b> <b><code>&nbsp;&nbsp;6125🍴</code></b> [pingcap/tidb](https://github.com/pingcap/tidb)) | 兼容 MySQL 协议的分布式 HTAP 数据库。 |
| <b><code>&nbsp;31353⭐</code></b> <b><code>&nbsp;&nbsp;3699🍴</code></b> [influxdata/influxdb](https://github.com/influxdata/influxdb)) | 经典的时序数据库项目。 |
| <b><code>&nbsp;21639⭐</code></b> <b><code>&nbsp;&nbsp;1584🍴</code></b> [dgraph-io/dgraph](https://github.com/dgraph-io/dgraph)) | 面向关联查询场景的分布式图数据库。 |
| <b><code>&nbsp;16938⭐</code></b> <b><code>&nbsp;&nbsp;3141🍴</code></b> [ipfs/kubo](https://github.com/ipfs/kubo)) | IPFS 的 Go 实现。 |
| <b><code>&nbsp;&nbsp;&nbsp;&nbsp;19⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2🍴</code></b> [chrislusf/seaweedfs](https://github.com/chrislusf/seaweedfs)) | 高性能分布式文件系统，覆盖对象、文件和块存储。 |
| <b><code>&nbsp;&nbsp;2772⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;451🍴</code></b> [XiaoMi/Gaea](https://github.com/XiaoMi/Gaea)) | 小米开源的 MySQL 中间件，聚焦分库分表与代理能力。 |
| <b><code>&nbsp;&nbsp;&nbsp;640⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;114🍴</code></b> [mediocregopher/radix](https://github.com/mediocregopher/radix)) | 设计简洁的 Go Redis 客户端。 |
| <b><code>&nbsp;&nbsp;7471⭐</code></b> <b><code>&nbsp;&nbsp;1155🍴</code></b> [olivere/elastic](https://github.com/olivere/elastic)) | Go 生态里长期被广泛使用的 Elasticsearch client。 |

## 可观测性

指标、图表、告警与运行状态检查

| 项目 | 简介 |
| --- | --- |
| <b><code>&nbsp;72509⭐</code></b> <b><code>&nbsp;13509🍴</code></b> [grafana/grafana](https://github.com/grafana/grafana)) | 最常见的可观测性可视化平台之一。 |
| <b><code>&nbsp;63067⭐</code></b> <b><code>&nbsp;10212🍴</code></b> [prometheus/prometheus](https://github.com/prometheus/prometheus)) | 事实标准级的监控与时序指标系统。 |
| <b><code>&nbsp;&nbsp;2372⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;487🍴</code></b> [influxdata/kapacitor](https://github.com/influxdata/kapacitor)) | InfluxData 的实时计算、告警与监控处理组件。 |
| <b><code>&nbsp;&nbsp;3458⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;248🍴</code></b> [sourcegraph/checkup](https://github.com/sourcegraph/checkup)) | 分布式健康检查工具，适合做站点和服务可用性探测。 |
| <b><code>&nbsp;&nbsp;2191⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;158🍴</code></b> [rapidloop/rtop](https://github.com/rapidloop/rtop)) | 基于 SSH 的轻量级远程服务器监控工具。 |

## 网络与安全

网关、负载均衡、代理、流量调试与网络工具

| 项目 | 简介 |
| --- | --- |
| <b><code>&nbsp;62057⭐</code></b> <b><code>&nbsp;&nbsp;5857🍴</code></b> [traefik/traefik](https://github.com/traefik/traefik)) | 云原生场景里广泛使用的反向代理和负载均衡器。 |
| <b><code>&nbsp;&nbsp;5679⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;508🍴</code></b> [google/seesaw](https://github.com/google/seesaw)) | Google 开源的 Linux 负载均衡系统。 |
| <b><code>&nbsp;&nbsp;&nbsp;796⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;199🍴</code></b> [jpillora/go-tcp-proxy](https://github.com/jpillora/go-tcp-proxy)) | 实现简单、非常适合学习 TCP 代理原理。 |
| <b><code>&nbsp;19252⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;84🍴</code></b> [probelabs/goreplay](https://github.com/probelabs/goreplay)) | 把线上 HTTP 流量复制回测试环境的经典工具。 |
| <b><code>&nbsp;&nbsp;&nbsp;277⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;89🍴</code></b> [hidu/pproxy](https://github.com/hidu/pproxy)) | HTTP 抓包代理和调试工具。 |
| <b><code>&nbsp;15258⭐</code></b> <b><code>&nbsp;11130🍴</code></b> [getlantern/lantern](https://github.com/getlantern/lantern)) | 长期维护的网络代理项目，可参考跨平台网络客户端设计。 |

## Web 开发与应用

Web 框架、服务端组件与实时交互能力

| 项目 | 简介 |
| --- | --- |
| <b><code>&nbsp;88200⭐</code></b> <b><code>&nbsp;&nbsp;8568🍴</code></b> [gin-gonic/gin](https://github.com/gin-gonic/gin)) | Go Web 框架里最常见的高性能选择。 |
| <b><code>&nbsp;32210⭐</code></b> <b><code>&nbsp;&nbsp;2311🍴</code></b> [labstack/echo](https://github.com/labstack/echo)) | API 开发体验成熟的高性能 Web 框架。 |
| <b><code>&nbsp;32427⭐</code></b> <b><code>&nbsp;&nbsp;5613🍴</code></b> [beego/beego](https://github.com/beego/beego)) | 老牌但仍在维护的全功能 Go Web 框架。 |
| <b><code>&nbsp;13251⭐</code></b> <b><code>&nbsp;&nbsp;1367🍴</code></b> [revel/revel](https://github.com/revel/revel)) | 偏完整栈思路的 Go Web 框架。 |
| <b><code>&nbsp;25618⭐</code></b> <b><code>&nbsp;&nbsp;2455🍴</code></b> [kataras/iris](https://github.com/kataras/iris)) | 强调性能和完整生态的 Go Web 框架。 |
| <b><code>&nbsp;&nbsp;3556⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;285🍴</code></b> [go-macaron/macaron](https://github.com/go-macaron/macaron)) | 模块化风格明显的 Go Web 框架。 |
| <b><code>&nbsp;&nbsp;1596⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;249🍴</code></b> [andeya/faygo](https://github.com/andeya/faygo)) | 面向 API 场景的 Go Web 框架，带参数绑定和文档生成。 |
| <b><code>&nbsp;&nbsp;4067⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;374🍴</code></b> [olahol/melody](https://github.com/olahol/melody)) | 基于 gorilla/websocket 的轻量级 WebSocket 框架。 |
| <b><code>&nbsp;23287⭐</code></b> <b><code>&nbsp;&nbsp;1818🍴</code></b> [valyala/fasthttp](https://github.com/valyala/fasthttp)) | Go 里非常有代表性的高性能 HTTP 实现。 |
| <b><code>&nbsp;&nbsp;3729⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;538🍴</code></b> [tus/tusd](https://github.com/tus/tusd)) | 断点续传文件上传服务端实现。 |
| <b><code>&nbsp;35683⭐</code></b> <b><code>&nbsp;&nbsp;8390🍴</code></b> [mattermost/mattermost](https://github.com/mattermost/mattermost)) | 大型 Go Web 应用的代表项目, 适合看真实业务系统的工程组织方式。 |

## 数据处理与机器学习

ML、NLP、爬虫与数据处理

| 项目 | 简介 |
| --- | --- |
| <b><code>&nbsp;&nbsp;5913⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;449🍴</code></b> [gorgonia/gorgonia](https://github.com/gorgonia/gorgonia)) | Go 生态里最有代表性的深度学习与张量计算项目之一。 |
| <b><code>&nbsp;&nbsp;1616⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;135🍴</code></b> [cdipaolo/goml](https://github.com/cdipaolo/goml)) | 提供在线学习、聚类和回归等算法实现。 |
| <b><code>&nbsp;&nbsp;9450⭐</code></b> <b><code>&nbsp;&nbsp;1180🍴</code></b> [sjwhitworth/golearn](https://github.com/sjwhitworth/golearn)) | 更偏传统机器学习流程的 Go 库。 |
| <b><code>&nbsp;&nbsp;7602⭐</code></b> <b><code>&nbsp;&nbsp;1686🍴</code></b> [andeya/pholcus](https://github.com/andeya/pholcus)) | Go 编写的分布式爬虫框架。 |
| <b><code>&nbsp;&nbsp;2622⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;306🍴</code></b> [yanyiwu/gojieba](https://github.com/yanyiwu/gojieba)) | 结巴中文分词的 Go 版本。 |
| <b><code>&nbsp;&nbsp;3556⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;292🍴</code></b> [chrislusf/gleam](https://github.com/chrislusf/gleam)) | Go 风格的数据处理和分布式计算框架。 |

## 开发者工具与基础库

开发效率、测试、终端 UI 和核心基础库

| 项目 | 简介 |
| --- | --- |
| <b><code>&nbsp;86922⭐</code></b> <b><code>&nbsp;&nbsp;8197🍴</code></b> [gohugoio/hugo](https://github.com/gohugoio/hugo)) | 最有代表性的 Go 静态站点生成器。 |
| <b><code>&nbsp;22820⭐</code></b> <b><code>&nbsp;&nbsp;4645🍴</code></b> [grpc/grpc-go](https://github.com/grpc/grpc-go)) | gRPC 的 Go 官方实现。 |
| <b><code>&nbsp;19804⭐</code></b> <b><code>&nbsp;&nbsp;1286🍴</code></b> [rakyll/hey](https://github.com/rakyll/hey)) | 轻量级压力测试工具。 |
| <b><code>&nbsp;&nbsp;7755⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;988🍴</code></b> [visualfc/liteide](https://github.com/visualfc/liteide)) | 跨平台的 Go IDE。 |
| <b><code>&nbsp;&nbsp;1498⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;195🍴</code></b> [mailslurper/mailslurper](https://github.com/mailslurper/mailslurper)) | 本地开发非常实用的测试 SMTP 服务器。 |
| <b><code>&nbsp;13524⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;817🍴</code></b> [gizak/termui](https://github.com/gizak/termui)) | 在终端里构建可视化面板的 Go UI 库。 |
| <b><code>&nbsp;&nbsp;6140⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;815🍴</code></b> [golang/mobile](https://github.com/golang/mobile)) | Go 官方维护的移动端开发工具链。 |
| <b><code>&nbsp;&nbsp;5817⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;595🍴</code></b> [hound-search/hound](https://github.com/hound-search/hound)) | 适合自建的代码搜索工具。 |

## 区块链

仍在维护、影响力最大的 Go 区块链项目

| 项目 | 简介 |
| --- | --- |
| <b><code>&nbsp;50892⭐</code></b> <b><code>&nbsp;21827🍴</code></b> [ethereum/go-ethereum](https://github.com/ethereum/go-ethereum)) | 以太坊客户端 geth 的官方实现。 |
| <b><code>&nbsp;16596⭐</code></b> <b><code>&nbsp;&nbsp;9110🍴</code></b> [hyperledger/fabric](https://github.com/hyperledger/fabric)) | 企业级联盟链平台的代表项目。 |

## 维护说明

目录已经去掉旧版 README 中的重复收录、过时仓库和“其它”大杂烩分类。后续如果继续扩展，建议优先更新 [projects.json](projects.json)，再运行 `go run ./tools/generate_readme.go` 同步生成 README。

## Source
<b><code>&nbsp;11413⭐</code></b> <b><code>&nbsp;&nbsp;2201🍴</code></b> [hackstoic/golang-open-source-projects](https://github.com/hackstoic/golang-open-source-projects))