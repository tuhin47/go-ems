
**Upload your config to Consul KV** (this is what your Go app expects):

```sh
curl --request PUT --data-binary @docs/consul-config.json \
  http://localhost:8500/v1/kv/event-management
```