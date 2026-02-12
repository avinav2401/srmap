![alt text](frontend/public/images/og.png)

# SrmAP
### Better way to manage your academics.

View, predict, and strategize your success.

> This monorepo contains both the frontend (Next.js) and backend (Go) for the SrmAP application.
> 
> ---
> 
> ## Monorepo Structure
> 
> ```
> srmap/
> ├── frontend/          # Next.js frontend application
> ├── backend/           # Go backend API
> ├── .env.example       # Environment variables template
> ├── package.json
> ├── compose.yaml
> └── README.md
> ```


### Prerequisites

- [Bun](https://bun.sh/) (>=1.2.0)
- [Go](https://golang.org/) (>=1.23.0)
- [Docker](https://docker.com/) (optional, for containerized deployment)

### Setup

1. **Clone the repository:**

   ```bash
   git clone --recurse-submodules <repository-url> srmap
   cd srmap
   ```

2. **Install dependencies:**

   ```bash
   # Install the run script
   bun install

   # Install all dependencies
   bun run install:all
   ```

3. **Environment Setup:**
Copy from `.env.example` and paste it in the root directory

```bash
# Shared Configuration
SUPABASE_URL="your_supabase_url"
SUPABASE_KEY="your_supabase_anon_key"
VALIDATION_KEY="your_validation_key"

# Frontend Specific (Mostly autofilled by the variables)
NEXT_PUBLIC_URL="http://localhost:8080"
NEXT_PUBLIC_SUPABASE_URL="${SUPABASE_URL}"
NEXT_PUBLIC_SERVICE_KEY="${SUPABASE_KEY}"
NEXT_PUBLIC_VALIDATION_KEY="${VALIDATION_KEY}"

# Backend Specific
ENCRYPTION_KEY="your_encryption_key"
URL="http://localhost:3000,http://localhost:0243"
```


> [!TIP]
> Generate secure keys for `VALIDATION_KEY` and `ENCRYPTION_KEY`.
>
> **For Linux, macOS, or Windows with Git Bash/WSL:**
>
> ```bash
> openssl rand -hex 32
> ```
>
> **For Windows with PowerShell:**
>
> ```powershell
> [BitConverter]::ToString((New-Object Security.Cryptography.RNGCryptoServiceProvider).GetBytes(32)).Replace("-", "").ToLower()
> ```

### Development

#### Run both services:

```bash
# Frontend (http://localhost:0243)
bun run dev:frontend

# Backend (http://localhost:8080)
bun run dev:backend

# Run the app as a whole
bun run dev
```

### Production Build

```bash
# Build both services as a whole
bun run build

# Build individually
bun run build:frontend
bun run build:backend
```

### Docker Deployment

```bash
# Build and run with Docker Compose
bun run docker:build
bun run docker:up

# Stop services
bun run docker:down
```


> [!WARNING]
> We will **NOT** take account for anything caused by your self-hosted instance


## Why Choose SrmAP?

- **Mobile-First Approach:** Built for mobile devices, Optimized for desktop and tablet devices.
- **Open Source:** Transparent and community-driven.
- **Massive Community**: Used by 16k+ students every month.
- **Timetable Generation:** Creates a full timetable based on your class schedule.
- **Attendance Prediction:** Predicts the percent based on your expected leave days
- **Safe and Secure:** Built with privacy and security in mind.
- **No Bloat:** Streamlined and efficient, with no unnecessary bloatware.

### The Idea Behind SrmAP

This project was intended to show the timetable and attendance. but it grew and scaled to a full-on replacement to SRM Academia. We made sure to use the web-standards and the best-in-class approaches to make sure our service is `fast`, `easy-to-use` and `easy on eyes`.

## How it Works

SrmAP integrates directly with your academic data to provide:
- **Real-time Attendance Tracking**: Never miss a class or fall below the margin.
- **Smart Analytics**: GPA prediction and performance insights.
- **Seamless Scheduling**: Automated timetable and calendar events.


---

## Contributing

Thinking about contributing to this project?

👉 **[Read our Contributing Guide](CONTRIBUTING.md)** to get started!

---

## [License](https://creativecommons.org/licenses/by-nc-nd/4.0/)

### You are free to:

- **Share:** Copy and redistribute the material in any medium or format.

### Under the following terms:

- **Attribution:** You must give appropriate credit, provide a link to the license, and indicate if changes were made. You may do so in any reasonable manner, but not in any way that suggests the licensor endorses you or your use.
- **NonCommercial:** You may not use the material for commercial purposes.
- **NoDerivatives:** If you remix, transform, or build upon the material, you may not distribute the modified material.
