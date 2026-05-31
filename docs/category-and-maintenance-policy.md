# Category and Maintenance Policy

[Chinese version](分类与维护策略.md)

## Goals of This Refactor

The old catalog had four main issues:

- Category granularity was inconsistent, and the boundaries between `Web Tools`, `Web Frameworks`, and `Other` were unclear.
- The same project could appear in multiple categories, which made lookup harder.
- The catalog mixed together "still worth recommending" and "historically appeared in the Go ecosystem".
- AI / Agent projects were barely covered, which no longer reflects current areas of interest in the Go ecosystem.

This refresh narrows the catalog into 10 topic categories and limits each project to a single entry.

## New Categories

| Category | Focus |
| --- | --- |
| AI / Agent | LLM application frameworks, MCP, model runtimes, vector search |
| Cloud Native and Containers | Container runtimes, orchestration, registries, cluster platforms |
| Service Governance and Platform Engineering | PaaS, service governance, CI/CD, messaging, async jobs |
| Data Storage and Search | Databases, distributed storage, search, data access ecosystems |
| Observability | Metrics, dashboards, alerting, health checks |
| Networking and Security | Gateways, proxies, load balancing, traffic debugging |
| Web Development and Applications | Web frameworks, WebSocket, HTTP components |
| Data Processing and Machine Learning | ML, NLP, crawlers, data processing |
| Developer Tools and Core Libraries | Core libraries, testing tools, terminal UI, IDEs |
| Blockchain | Maintained and sufficiently influential Go blockchain projects |

## Selection Criteria

Maintenance status review date: `2026-03-06`

- Prefer projects that could still be confirmed as maintained.
- Remove archived repositories directly.
- Remove long-inactive projects by default when a more suitable alternative now exists.
- Emphasize projects that are still worth learning from or selecting today, rather than projects that merely appeared historically.
- Do not include the same project in multiple categories.

## Update Recommendations

For future maintenance, use this order:

1. Update [projects.json](../projects.json) first.
2. Run `go run ./tools/generate_readme.go`.
3. If there are large migrations or removals, update the [removal and migration log](removal-and-migration-log.md) as well.
4. Keep [README_EN.md](../README_EN.md) aligned with the generated Chinese README.
