# gonfservice
Simple sample service API to send e-mail and other types of notifications using golang

## Libs

- https://github.com/go-gomail/gomail

## Development

Use `realize` to enable auto-reload during development:

```shell
> realize start
```

## Functionalities

- [x] Send email throught HTTP endpoint
- [ ] Enqueue messages to send later
- [ ] Check sending status
- [ ] Create and use message templates (Markdown)