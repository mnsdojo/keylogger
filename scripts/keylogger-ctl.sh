#!/bin/bash

SERVICE_NAME="keylogger"
LOG_DIR="$HOME/keylogger/logs"

# Ensure the logs directory exists
mkdir -p "$LOG_DIR"

case $1 in
start)
    echo "Starting $SERVICE_NAME service..."
    sudo systemctl start $SERVICE_NAME
    ;;
stop)
    echo "Stopping $SERVICE_NAME service..."
    sudo systemctl stop $SERVICE_NAME
    ;;
restart)
    echo "Restarting $SERVICE_NAME service..."
    sudo systemctl restart $SERVICE_NAME
    ;;
status)
    echo "Checking status of $SERVICE_NAME service..."
    sudo systemctl status $SERVICE_NAME
    ;;
logs)
    if [ -f "$LOG_DIR/keylogger.log" ]; then
        echo "Tailing logs for $SERVICE_NAME..."
        tail -f "$LOG_DIR/keylogger.log"
    else
        echo "Log file not found at $LOG_DIR/keylogger.log"
    fi
    ;;
enable)
    echo "Enabling $SERVICE_NAME to start on boot..."
    sudo systemctl enable $SERVICE_NAME
    ;;
disable)
    echo "Disabling $SERVICE_NAME from starting on boot..."
    sudo systemctl disable $SERVICE_NAME
    ;;
*)
    echo "Usage: $0 {start|stop|restart|status|logs|enable|disable}"
    exit 1
    ;;
esac

exit 0
