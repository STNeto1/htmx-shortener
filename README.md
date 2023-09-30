## HTMX Link Shortener

### How to deploy

```bash
$ fly launch --no-deploy                                                                                                                      10:00:43
Creating app in /Users/nyx/Apps/htmx-shortener
Scanning source code
Detected a Dockerfile app
? Choose an app name (leave blank to generate one): htmx-shortener
? Select Organization: ... (personal)
Some regions require a paid plan (bom, fra, maa).
See https://fly.io/plans to set up a plan.

? Choose a region for deployment: Rio de Janeiro, Brazil (gig)
App will use 'gig' region as primary

Created app 'htmx-shortener' in organization 'personal'
Admin URL: https://fly.io/apps/htmx-shortener
Hostname: htmx-shortener.fly.dev
? Create .dockerignore from 1 .gitignore files? Yes
Created /Users/nyx/Apps/htmx-shortener/.dockerignore from 1 .gitignore files.
Wrote config file fly.toml
Validating /Users/nyx/Apps/htmx-shortener/fly.toml
Platform: machines
✓ Configuration is valid
Your app is ready! Deploy with `flyctl deploy`
```

```bash
# replace gig with the same region as the fly app
$ turso db create shortener --location gig
$ turso db show shortener                                                                                                                     10:02:45
Name:           shortener
URL:            libsql://....turso.io
ID:             ...
Group:          default
Version:        0.21.2
Locations:      gig
Size:           8.2 kB

Database Instances:
NAME     TYPE        LOCATION
gig      primary     gig

$ turso db tokens create shortener                                                                                                            10:02:07
[[TOKEN]]
```

```bash
$ fly secrets set PORT="3000" BASE_URL="https://htmx-shortener.fly.dev" DATABASE_URL="libsql://....turso.io?authToken=tokenFromBefore"
$ fly deploy
```

Now accessing the url from fly should display the main page

---

### What's next?

- Add authentication
- Index database
- Analytics?
- Add testing
