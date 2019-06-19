#! /bin/bash
export LC_ALL=en_US.UTF-8
export LANG=en_US.UTF-8

ROOT=`cd $(dirname $0); pwd`
BIN=${ROOT}/admission-webhook
LOGDIR=${ROOT}/logs/

mkdir -p ${LOGDIR}
process="admission-webhook"

#使用pgrep判断程序是否执行
function check_pid(){
    run_pid=$(pgrep $process)
        echo $run_pid
}
pid=$(check_pid)

function status(){
    if [ "x_$pid" != "x_" ]; then
        echo "$process running with pid: $pid"
    else
        echo "ERROR: admission-webhook may not running!"
    fi
}

start() {
    echo -n "starting admission-webhook... "
    exec ${ROOT}/admission-webhook --tls-cert-file=/etc/kubernetes/pki/admission-webhook-tls.crt \
--tls-private-key-file=/etc/kubernetes/pki/admission-webhook-tls.key
}

stop() {
    echo -n "stopping admission-webhook... "
    kill `cat ${pid}`
    return 0
}

restart() {
    echo -n "restarting admission-webhook... "
    stop
    start
    echo "finished, plz check by urself"
}

case "$1" in
start)
    start
    ;;

stop)
    stop
    ;;

restart)
    restart
    ;;
*)
    echo "Usage: $0 {start|stop|restart}"
    exit 1
    esac
