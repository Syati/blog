# blog

https://blog.syati.info

## Setup

### Requirements
- [mise](https://mise.jdx.dev/) - Tool version manager
- Go 1.23+ (automatically installed via mise)

### Installation

```bash
# Install mise (if not already installed)
curl https://mise.run | sh

# Install Go via mise
mise install

# Install Hugo
mise run install-hugo
```

## Development

### Build search index
```bash
mise run build-index
```

### Build site
```bash
mise run build
```

### Start development server
```bash
mise run serve
```

### Deploy to GitHub Pages
```bash
mise run deploy
# or
./deploy.sh
```

## Tech Stack

- **Hugo v0.160+** - Static site generator
- **Go 1.23** - Search index generation
- **mise** - Development environment management
