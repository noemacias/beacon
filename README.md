![Beacon Gopher Logo](images/beacon-gopher.png)

# Beacon

A lightweight, configurable personal landing page that lets you showcase your profile, links, resume, and contact information with multiple built-in themes.

# Complete Configuration Example

```yaml
server:
  addr: :8080
  tls:
    enabled: false
    cert: /etc/ssl/certs/server.crt
    key: /etc/ssl/private/server.key

site:
  title: John Doe
  subtitle: Senior Software Engineer
  description: Building scalable systems and cloud-native applications.
  avatar:
    url: https://example.com/avatar.jpg
    name: John Doe
  initials: JD
  static_dir: ./static

theme:
  name: dark.html

links:
  - name: GitHub
    icon: GH
    url: https://github.com/johndoe
    description: Browse my open-source projects

  - name: LinkedIn
    icon: in
    url: https://linkedin.com/in/johndoe
    description: Connect with me professionally

  - name: Resume
    icon: CV
    url: /static/resume.pdf
    description: Download my resume

  - name: Email
    icon: "@"
    url: mailto:john.doe@example.com
    description: Send me an email

  - name: Blog
    icon: BL
    url: https://blog.johndoe.dev
    description: Articles on software engineering and DevOps
```

---

# Configuration Options

## Server

| Field                | Default | Description                                 |
| -------------------- | ------- | ------------------------------------------- |
| `server.addr`        | `:8080` | Address and port the web server listens on. |
| `server.tls.enabled` | `false` | Enable HTTPS/TLS support.                   |
| `server.tls.cert`    | `""`    | Path to the TLS certificate file.           |
| `server.tls.key`     | `""`    | Path to the TLS private key file.           |

## Site

| Field              | Default    | Description                                             |
| ------------------ | ---------- | ------------------------------------------------------- |
| `site.title`       | `""`       | Main title displayed on the homepage.                   |
| `site.subtitle`    | `""`       | Secondary title displayed beneath the main title.       |
| `site.description` | `""`       | Short description or tagline.                           |
| `site.avatar.url`  | `""`       | URL of the avatar image.                                |
| `site.avatar.name` | `""`       | Alt text for the avatar image.                          |
| `site.initials`    | `""`       | Fallback initials displayed when no avatar is provided. |
| `site.static_dir`  | `./static` | Directory used to serve static files.                   |

## Theme

| Field        | Default     | Description              |
| ------------ | ----------- | ------------------------ |
| `theme.name` | `dark.html` | Theme template filename. |

## Links

Each entry in `links` supports:

| Field         | Required | Description                              |
| ------------- | -------- | ---------------------------------------- |
| `name`        | Yes      | Display name for the link.               |
| `icon`        | No       | Short icon text or symbol.               |
| `url`         | Yes      | Destination URL.                         |
| `description` | No       | Additional text displayed with the link. |

Supported URL formats:

```yaml
# External URL
url: https://github.com/johndoe

# Email
url: mailto:john.doe@example.com

# Static file
url: /static/resume.pdf
```

## Available Themes

The application includes the following built-in themes:

| Theme           | Description                                                  |
| --------------- | ------------------------------------------------------------ |
| `dark.html`     | Default dark theme with a modern appearance.                 |
| `light.html`    | Light theme optimized for bright environments.               |
| `terminal.html` | Terminal-inspired theme with a retro command-line aesthetic. |

### Example

```yaml
theme:
  name: dark.html
```

### Available Values

```yaml
# Dark Theme (default)
theme:
  name: dark.html
```

```yaml
# Light Theme
theme:
  name: light.html
```

```yaml
# Terminal Theme
theme:
  name: terminal.html
```

> **Default:** `dark.html`
