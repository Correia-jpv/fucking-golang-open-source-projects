# Golang Open Source Projects

面向中文读者重新整理的 Go 开源项目目录。新版目录不再追求“尽可能全”，而是优先保留仍在维护、社区认知清晰、适合学习和选型的项目，并补充了 AI Agent 相关项目。

当前版本收录 **77** 个项目，分成 **10** 个主题；最近一次维护状态审阅时间为 **2026-03-06**。

- [English version](README_EN.md)
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
| AI / Agent | LLM 应用框架、MCP、模型运行时与向量能力 | 8 |
| 云原生与容器 | 容器运行时、编排、镜像仓库和集群平台 | 8 |
| 服务治理与平台工程 | PaaS、服务治理、CI/CD、消息与异步任务 | 12 |
| 数据存储与搜索 | 数据库、分布式存储、检索与数据访问生态 | 10 |
| 可观测性 | 指标、图表、告警与运行状态检查 | 6 |
| 网络与安全 | 网关、负载均衡、代理、流量调试与网络工具 | 6 |
| Web 开发与应用 | Web 框架、服务端组件与实时交互能力 | 11 |
| 数据处理与机器学习 | ML、NLP、爬虫与数据处理 | 6 |
| 开发者工具与基础库 | 开发效率、测试、终端 UI 和核心基础库 | 8 |
| 区块链 | 仍在维护、影响力最大的 Go 区块链项目 | 2 |

## AI / Agent

LLM 应用框架、MCP、模型运行时与向量能力

| 项目 | 简介 |
| --- | --- |
| <b><code>181764⭐</code></b> <b><code>&nbsp;18007🍴</code></b> [ollama/ollama](https://github.com/ollama/ollama)) | 本地运行、分发和管理大模型的 Go 运行时。 |
| <b><code>&nbsp;&nbsp;9696⭐</code></b> <b><code>&nbsp;&nbsp;1149🍴</code></b> [tmc/langchaingo](https://github.com/tmc/langchaingo)) | Go 版 LLM 应用框架，覆盖 prompt、tool calling、agent 和 RAG。 |
| <b><code>&nbsp;13166⭐</code></b> <b><code>&nbsp;&nbsp;1113🍴</code></b> [cloudwego/eino](https://github.com/cloudwego/eino)) | CloudWeGo 出品的 Go AI 应用框架，强调组件化编排和生产落地。 |
| <b><code>&nbsp;&nbsp;9143⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;887🍴</code></b> [mark3labs/mcp-go](https://github.com/mark3labs/mcp-go)) | 用 Go 构建 MCP client 和 server 的实用 SDK。 |
| <b><code>&nbsp;49273⭐</code></b> <b><code>&nbsp;&nbsp;4472🍴</code></b> [mudler/LocalAI](https://github.com/mudler/LocalAI)) | OpenAI 兼容的本地推理服务，适合私有化部署。 |
| <b><code>&nbsp;&nbsp;1980⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;292🍴</code></b> [mudler/LocalAGI](https://github.com/mudler/LocalAGI)) | 面向本地模型的 Agent 平台，强调工具调用和自治流程。 |
| <b><code>&nbsp;16850⭐</code></b> <b><code>&nbsp;&nbsp;1409🍴</code></b> [weaviate/weaviate](https://github.com/weaviate/weaviate)) | Go 编写的向量数据库，可用于 RAG、检索和 Agent memory。 |
| <b><code>&nbsp;&nbsp;&nbsp;536⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;46🍴</code></b> [pardnchiu/Agenvoy](https://github.com/pardnchiu/Agenvoy)) | Go 编写的 Agent 平台，提供 Py/Js 工具接口、错误记忆与自动修正能力。 |

## 云原生与容器

容器运行时、编排、镜像仓库和集群平台

| 项目 | 简介 |
| --- | --- |
| <b><code>&nbsp;72135⭐</code></b> <b><code>&nbsp;19235🍴</code></b> [moby/moby](https://github.com/moby/moby)) | Docker 引擎的上游项目，也是学习容器运行时实现的核心入口。 |
| <b><code>128015⭐</code></b> <b><code>&nbsp;45279🍴</code></b> [kubernetes/kubernetes](https://github.com/kubernetes/kubernetes)) | 事实标准级的容器编排平台。 |
| <b><code>&nbsp;29457⭐</code></b> <b><code>&nbsp;&nbsp;5362🍴</code></b> [goharbor/harbor](https://github.com/goharbor/harbor)) | 企业级 OCI 镜像仓库，带权限、审计和复制能力。 |
| <b><code>&nbsp;25938⭐</code></b> <b><code>&nbsp;&nbsp;3221🍴</code></b> [rancher/rancher](https://github.com/rancher/rancher)) | 面向多集群场景的 Kubernetes 管理平台。 |
| <b><code>&nbsp;11065⭐</code></b> <b><code>&nbsp;&nbsp;1217🍴</code></b> [quay/clair](https://github.com/quay/clair)) | 容器镜像漏洞分析与扫描服务。 |
| <b><code>&nbsp;&nbsp;3653⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;678🍴</code></b> [moby/swarmkit](https://github.com/moby/swarmkit)) | Docker Swarm 的核心编排组件，适合学习调度和集群编排。 |
| <b><code>&nbsp;&nbsp;4641⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;936🍴</code></b> [AliyunContainerService/pouch](https://github.com/AliyunContainerService/pouch)) | 阿里开源的容器引擎项目，聚焦更强的隔离与稳定性。 |
| <b><code>&nbsp;16974⭐</code></b> <b><code>&nbsp;&nbsp;2142🍴</code></b> [hashicorp/nomad](https://github.com/hashicorp/nomad)) | 轻量级工作负载编排器，适合对比 Kubernetes 的另一条路线。 |

## 服务治理与平台工程

PaaS、服务治理、CI/CD、消息与异步任务

| 项目 | 简介 |
| --- | --- |
| <b><code>&nbsp;&nbsp;5312⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;552🍴</code></b> [tsuru/tsuru](https://github.com/tsuru/tsuru)) | 成熟的开源 PaaS，适合学习应用平台抽象。 |
| <b><code>&nbsp;&nbsp;6267⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;878🍴</code></b> [goodrain/rainbond](https://github.com/goodrain/rainbond)) | 以应用为中心的云原生平台，覆盖交付、运维和微服务治理。 |
| <b><code>&nbsp;38441⭐</code></b> <b><code>&nbsp;&nbsp;3409🍴</code></b> [harness/harness](https://github.com/harness/harness)) | Drone 已并入 Harness 生态后，新的 CI/CD 与开发者平台入口。 |
| <b><code>&nbsp;20946⭐</code></b> <b><code>&nbsp;&nbsp;2162🍴</code></b> [gravitational/teleport](https://github.com/gravitational/teleport)) | 基于零信任模型的远程访问与基础设施入口。 |
| <b><code>&nbsp;38408⭐</code></b> <b><code>&nbsp;&nbsp;8374🍴</code></b> [istio/istio](https://github.com/istio/istio)) | 服务网格代表项目，覆盖流量治理、安全和可观测性。 |
| <b><code>&nbsp;&nbsp;&nbsp;&nbsp;11⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;4🍴</code></b> [uber/jaeger](https://github.com/uber/jaeger)) | 分布式追踪系统，适合与 OpenTelemetry 一起理解链路追踪。 |
| <b><code>&nbsp;27424⭐</code></b> <b><code>&nbsp;&nbsp;2438🍴</code></b> [go-kit/kit](https://github.com/go-kit/kit)) | Go 微服务开发工具箱，强调可观测性和可测试性。 |
| <b><code>&nbsp;&nbsp;6110⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;582🍴</code></b> [goadesign/goa](https://github.com/goadesign/goa)) | 设计优先的 Go 服务开发框架。 |
| <b><code>&nbsp;10834⭐</code></b> <b><code>&nbsp;&nbsp;1166🍴</code></b> [TykTechnologies/tyk](https://github.com/TykTechnologies/tyk)) | 成熟的开源 API Gateway。 |
| <b><code>&nbsp;23079⭐</code></b> <b><code>&nbsp;&nbsp;2425🍴</code></b> [micro/go-micro](https://github.com/micro/go-micro)) | Go 微服务框架，适合研究服务抽象与插件化扩展。 |
| <b><code>&nbsp;25780⭐</code></b> <b><code>&nbsp;&nbsp;2886🍴</code></b> [nsqio/nsq](https://github.com/nsqio/nsq)) | 经典的实时分布式消息平台。 |
| <b><code>&nbsp;&nbsp;7972⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;933🍴</code></b> [RichardKnop/machinery](https://github.com/RichardKnop/machinery)) | Go 异步任务队列，适合替代 Celery 的思路参考。 |

## 数据存储与搜索

数据库、分布式存储、检索与数据访问生态

| 项目 | 简介 |
| --- | --- |
| <b><code>&nbsp;32508⭐</code></b> <b><code>&nbsp;&nbsp;4116🍴</code></b> [cockroachdb/cockroach](https://github.com/cockroachdb/cockroach)) | 分布式 SQL 数据库，强调强一致与弹性扩展。 |
| <b><code>&nbsp;21355⭐</code></b> <b><code>&nbsp;&nbsp;2407🍴</code></b> [vitessio/vitess](https://github.com/vitessio/vitess)) | YouTube 开源的 MySQL 水平扩展方案。 |
| <b><code>&nbsp;40591⭐</code></b> <b><code>&nbsp;&nbsp;6247🍴</code></b> [pingcap/tidb](https://github.com/pingcap/tidb)) | 兼容 MySQL 协议的分布式 HTAP 数据库。 |
| <b><code>&nbsp;31759⭐</code></b> <b><code>&nbsp;&nbsp;3716🍴</code></b> [influxdata/influxdb](https://github.com/influxdata/influxdb)) | 经典的时序数据库项目。 |
| <b><code>&nbsp;21804⭐</code></b> <b><code>&nbsp;&nbsp;1605🍴</code></b> [dgraph-io/dgraph](https://github.com/dgraph-io/dgraph)) | 面向关联查询场景的分布式图数据库。 |
| <b><code>&nbsp;17142⭐</code></b> <b><code>&nbsp;&nbsp;3173🍴</code></b> [ipfs/kubo](https://github.com/ipfs/kubo)) | IPFS 的 Go 实现。 |
| <b><code>&nbsp;&nbsp;&nbsp;&nbsp;41⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;7🍴</code></b> [chrislusf/seaweedfs](https://github.com/chrislusf/seaweedfs)) | 高性能分布式文件系统，覆盖对象、文件和块存储。 |
| <b><code>&nbsp;&nbsp;2765⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;449🍴</code></b> [XiaoMi/Gaea](https://github.com/XiaoMi/Gaea)) | 小米开源的 MySQL 中间件，聚焦分库分表与代理能力。 |
| <b><code>&nbsp;&nbsp;&nbsp;637⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;112🍴</code></b> [mediocregopher/radix](https://github.com/mediocregopher/radix)) | 设计简洁的 Go Redis 客户端。 |
| <b><code>&nbsp;&nbsp;7443⭐</code></b> <b><code>&nbsp;&nbsp;1142🍴</code></b> [olivere/elastic](https://github.com/olivere/elastic)) | Go 生态里长期被广泛使用的 Elasticsearch client。 |

## 可观测性

指标、图表、告警与运行状态检查

| 项目 | 简介 |
| --- | --- |
| <b><code>&nbsp;76926⭐</code></b> <b><code>&nbsp;14790🍴</code></b> [grafana/grafana](https://github.com/grafana/grafana)) | 最常见的可观测性可视化平台之一。 |
| <b><code>&nbsp;66248⭐</code></b> <b><code>&nbsp;10861🍴</code></b> [prometheus/prometheus](https://github.com/prometheus/prometheus)) | 事实标准级的监控与时序指标系统。 |
| <b><code>&nbsp;&nbsp;2376⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;478🍴</code></b> [influxdata/kapacitor](https://github.com/influxdata/kapacitor)) | InfluxData 的实时计算、告警与监控处理组件。 |
| <b><code>&nbsp;&nbsp;3459⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;246🍴</code></b> [sourcegraph/checkup](https://github.com/sourcegraph/checkup)) | 分布式健康检查工具，适合做站点和服务可用性探测。 |
| <b><code>&nbsp;&nbsp;2189⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;156🍴</code></b> [rapidloop/rtop](https://github.com/rapidloop/rtop)) | 基于 SSH 的轻量级远程服务器监控工具。 |
| <b><code>&nbsp;&nbsp;&nbsp;139⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;129🍴</code></b> [kubestellar/console](https://github.com/kubestellar/console)) | AI 驱动的多集群 Kubernetes 仪表盘，支持实时可观测性和 30+ CNCF 项目集成。CNCF Sandbox 项目。 |

## 网络与安全

网关、负载均衡、代理、流量调试与网络工具

| 项目 | 简介 |
| --- | --- |
| <b><code>&nbsp;64972⭐</code></b> <b><code>&nbsp;&nbsp;6206🍴</code></b> [traefik/traefik](https://github.com/traefik/traefik)) | 云原生场景里广泛使用的反向代理和负载均衡器。 |
| <b><code>&nbsp;&nbsp;5670⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;505🍴</code></b> [google/seesaw](https://github.com/google/seesaw)) | Google 开源的 Linux 负载均衡系统。 |
| <b><code>&nbsp;&nbsp;&nbsp;802⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;196🍴</code></b> [jpillora/go-tcp-proxy](https://github.com/jpillora/go-tcp-proxy)) | 实现简单、非常适合学习 TCP 代理原理。 |
| <b><code>&nbsp;19324⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;99🍴</code></b> [probelabs/goreplay](https://github.com/probelabs/goreplay)) | 把线上 HTTP 流量复制回测试环境的经典工具。 |
| <b><code>&nbsp;&nbsp;&nbsp;277⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;&nbsp;88🍴</code></b> [hidu/pproxy](https://github.com/hidu/pproxy)) | HTTP 抓包代理和调试工具。 |
| <b><code>&nbsp;16044⭐</code></b> <b><code>&nbsp;11044🍴</code></b> [getlantern/lantern](https://github.com/getlantern/lantern)) | 长期维护的网络代理项目，可参考跨平台网络客户端设计。 |

## Web 开发与应用

Web 框架、服务端组件与实时交互能力

| 项目 | 简介 |
| --- | --- |
| <b><code>&nbsp;89261⭐</code></b> <b><code>&nbsp;&nbsp;8718🍴</code></b> [gin-gonic/gin](https://github.com/gin-gonic/gin)) | Go Web 框架里最常见的高性能选择。 |
| <b><code>&nbsp;32730⭐</code></b> <b><code>&nbsp;&nbsp;3307🍴</code></b> [labstack/echo](https://github.com/labstack/echo)) | API 开发体验成熟的高性能 Web 框架。 |
| <b><code>&nbsp;32429⭐</code></b> <b><code>&nbsp;&nbsp;5571🍴</code></b> [beego/beego](https://github.com/beego/beego)) | 老牌但仍在维护的全功能 Go Web 框架。 |
| <b><code>&nbsp;13215⭐</code></b> <b><code>&nbsp;&nbsp;1353🍴</code></b> [revel/revel](https://github.com/revel/revel)) | 偏完整栈思路的 Go Web 框架。 |
| <b><code>&nbsp;25567⭐</code></b> <b><code>&nbsp;&nbsp;2426🍴</code></b> [kataras/iris](https://github.com/kataras/iris)) | 强调性能和完整生态的 Go Web 框架。 |
| <b><code>&nbsp;&nbsp;3545⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;287🍴</code></b> [go-macaron/macaron](https://github.com/go-macaron/macaron)) | 模块化风格明显的 Go Web 框架。 |
| <b><code>&nbsp;&nbsp;1590⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;247🍴</code></b> [andeya/faygo](https://github.com/andeya/faygo)) | 面向 API 场景的 Go Web 框架，带参数绑定和文档生成。 |
| <b><code>&nbsp;&nbsp;4083⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;371🍴</code></b> [olahol/melody](https://github.com/olahol/melody)) | 基于 gorilla/websocket 的轻量级 WebSocket 框架。 |
| <b><code>&nbsp;23476⭐</code></b> <b><code>&nbsp;&nbsp;1869🍴</code></b> [valyala/fasthttp](https://github.com/valyala/fasthttp)) | Go 里非常有代表性的高性能 HTTP 实现。 |
| <b><code>&nbsp;&nbsp;3879⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;555🍴</code></b> [tus/tusd](https://github.com/tus/tusd)) | 断点续传文件上传服务端实现。 |
| <b><code>&nbsp;39192⭐</code></b> <b><code>&nbsp;&nbsp;9020🍴</code></b> [mattermost/mattermost](https://github.com/mattermost/mattermost)) | 大型 Go Web 应用的代表项目, 适合看真实业务系统的工程组织方式。 |

## 数据处理与机器学习

ML、NLP、爬虫与数据处理

| 项目 | 简介 |
| --- | --- |
| <b><code>&nbsp;&nbsp;5930⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;451🍴</code></b> [gorgonia/gorgonia](https://github.com/gorgonia/gorgonia)) | Go 生态里最有代表性的深度学习与张量计算项目之一。 |
| <b><code>&nbsp;&nbsp;1615⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;132🍴</code></b> [cdipaolo/goml](https://github.com/cdipaolo/goml)) | 提供在线学习、聚类和回归等算法实现。 |
| <b><code>&nbsp;&nbsp;9435⭐</code></b> <b><code>&nbsp;&nbsp;1164🍴</code></b> [sjwhitworth/golearn](https://github.com/sjwhitworth/golearn)) | 更偏传统机器学习流程的 Go 库。 |
| <b><code>&nbsp;&nbsp;7583⭐</code></b> <b><code>&nbsp;&nbsp;1665🍴</code></b> [andeya/pholcus](https://github.com/andeya/pholcus)) | Go 编写的分布式爬虫框架。 |
| <b><code>&nbsp;&nbsp;2645⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;305🍴</code></b> [yanyiwu/gojieba](https://github.com/yanyiwu/gojieba)) | 结巴中文分词的 Go 版本。 |
| <b><code>&nbsp;&nbsp;3566⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;292🍴</code></b> [chrislusf/gleam](https://github.com/chrislusf/gleam)) | Go 风格的数据处理和分布式计算框架。 |

## 开发者工具与基础库

开发效率、测试、终端 UI 和核心基础库

| 项目 | 简介 |
| --- | --- |
| <b><code>&nbsp;89955⭐</code></b> <b><code>&nbsp;&nbsp;8388🍴</code></b> [gohugoio/hugo](https://github.com/gohugoio/hugo)) | 最有代表性的 Go 静态站点生成器。 |
| <b><code>&nbsp;23075⭐</code></b> <b><code>&nbsp;&nbsp;4771🍴</code></b> [grpc/grpc-go](https://github.com/grpc/grpc-go)) | gRPC 的 Go 官方实现。 |
| <b><code>&nbsp;20286⭐</code></b> <b><code>&nbsp;&nbsp;1304🍴</code></b> [rakyll/hey](https://github.com/rakyll/hey)) | 轻量级压力测试工具。 |
| <b><code>&nbsp;&nbsp;7766⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;984🍴</code></b> [visualfc/liteide](https://github.com/visualfc/liteide)) | 跨平台的 Go IDE。 |
| <b><code>&nbsp;&nbsp;1510⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;196🍴</code></b> [mailslurper/mailslurper](https://github.com/mailslurper/mailslurper)) | 本地开发非常实用的测试 SMTP 服务器。 |
| <b><code>&nbsp;13590⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;820🍴</code></b> [gizak/termui](https://github.com/gizak/termui)) | 在终端里构建可视化面板的 Go UI 库。 |
| <b><code>&nbsp;&nbsp;6215⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;824🍴</code></b> [golang/mobile](https://github.com/golang/mobile)) | Go 官方维护的移动端开发工具链。 |
| <b><code>&nbsp;&nbsp;5883⭐</code></b> <b><code>&nbsp;&nbsp;&nbsp;600🍴</code></b> [hound-search/hound](https://github.com/hound-search/hound)) | 适合自建的代码搜索工具。 |

## 区块链

仍在维护、影响力最大的 Go 区块链项目

| 项目 | 简介 |
| --- | --- |
| <b><code>&nbsp;51371⭐</code></b> <b><code>&nbsp;22165🍴</code></b> [ethereum/go-ethereum](https://github.com/ethereum/go-ethereum)) | 以太坊客户端 geth 的官方实现。 |
| <b><code>&nbsp;16732⭐</code></b> <b><code>&nbsp;&nbsp;9105🍴</code></b> [hyperledger/fabric](https://github.com/hyperledger/fabric)) | 企业级联盟链平台的代表项目。 |

## 维护说明

目录已经去掉旧版 README 中的重复收录、过时仓库和“其它”大杂烩分类。后续如果继续扩展，建议优先更新 [projects.json](projects.json)，再运行 `go run ./tools/generate_readme.go` 同步生成 README。

## Source
<b><code>&nbsp;11564⭐</code></b> <b><code>&nbsp;&nbsp;2189🍴</code></b> [hackstoic/golang-open-source-projects](https://github.com/hackstoic/golang-open-source-projects))