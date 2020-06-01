# system-care

## Get

```
go get github.com/nguyenvanduocit/system-care
```

## Config

Create .system-care.yml on homedir (~/.system-care.yml)

```
token: 
org: 
bucket: 
server: 
```

run

```
system-care push
```

## LaunchAgents

Create file `~/Library/LaunchAgents/self.system-care.plist`

```
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
  <dict>
    <key>Label</key>
    <string>com.system-care</string>
    <key>WorkingDirectory</key>
    <string>HOME_DIR</string>
    <key>ProgramArguments</key>
    <array>
        <string>PATH_TO_SYSTEM_CARE_BIN</string>
        <string>push</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
    <key>KeepAlive</key>
    <true/>
    <key>StandardErrorPath</key>
    <string>PATH_TO_ERROR_FILE</string>
    <key>StandardOutPath</key>
    <string>PATH_TO_LOG_FILE</string>
  </dict>
</plist>
```

### Load

```
launchctl load ~/Library/LaunchAgents/self.system-care.plist
```

### Start

```
launchctl start com.system-care
```

### Check

```
launchctl list | grep system-care
```
