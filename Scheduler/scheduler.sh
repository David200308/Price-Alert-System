#!/bin/bash

ALERT_SCHEDULER_PATH="./Scheduler"
ALERT_DAILY_UPDATE_PATH="./daily_update"

ALERT_SCHEDULER_LOG="/var/log/alert_scheduler.log"
ALERT_DAILY_UPDATE_LOG="/var/log/alert_daily_update.log"

chmod +x "$ALERT_SCHEDULER_PATH"
chmod +x "$ALERT_DAILY_UPDATE_PATH"

crontab -l > crontab_backup_$(date +%F_%T) 2>/dev/null
echo "Existing crontab backed up."

CRON_JOBS=$(cat <<EOF
# Run ./Scheduler every 5 minutes
*/5 * * * * $ALERT_SCHEDULER_PATH >> $ALERT_SCHEDULER_LOG 2>&1

# Run ./daily_update at 00:00 every day
0 0 * * * $ALERT_DAILY_UPDATE_PATH >> $ALERT_DAILY_UPDATE_LOG 2>&1
EOF
)

(
  crontab -l 2>/dev/null
  echo "$CRON_JOBS"
) | crontab -

echo "Cron jobs added successfully."

echo "Updated crontab:"
crontab -l
