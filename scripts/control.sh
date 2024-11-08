#!/bin/bash

# 发生错误则退出
set -e


SCRIPT_DIR=$(dirname "$(readlink -f "$0")")  # 当前脚本所在目录
EXECUTABLE_FILE=""
OS=$(uname)  # 当前操作系统类型

echo "当前操作系统为${OS}"
if [ "${OS}" == "Darwin" ]; then
    echo "Running on macOS"
    EXECUTABLE_FILE="mac_main"
elif [ "${OS}" == "Linux" ]; then
    echo "Running on Linux"
    EXECUTABLE_FILE="linux_main"
else
    echo "Unsupported operating system: $OS"
    exit 1
fi

Ps(){
  echo "pgrep $EXECUTABLE_FILE"
  if pgrep $EXECUTABLE_FILE ; then
    echo "进程已启动"
  else
    echo "进程未启动"
  fi
}

Stop(){
  # 检查是否存在进程
  echo "pgrep $EXECUTABLE_FILE"
  if pgrep $EXECUTABLE_FILE > /dev/null; then
      # 如果存在则通过kill命令杀死进程
      echo "开始关闭"
      echo "pkill $EXECUTABLE_FILE"
      pkill $EXECUTABLE_FILE
      for ((i=0; i<10; ++i))
        do
          sleep 1
          if pgrep $EXECUTABLE_FILE > /dev/null; then
            echo -e ".\c"
          else
            echo 'Stop Success!'
            break;
          fi
        done

      if pgrep $EXECUTABLE_FILE > /dev/null; then
        echo "pkill -9 $EXECUTABLE_FILE"
        pkill -9 $EXECUTABLE_FILE
      fi
  else
      echo "进程已关闭"
  fi
}

Start(){
  echo "nohup $SCRIPT_DIR/$EXECUTABLE_FILE -config=$SCRIPT_DIR/config.yml -logpath=$SCRIPT_DIR/logs > /dev/null 2>&1 &"
  nohup $SCRIPT_DIR/$EXECUTABLE_FILE -config=$SCRIPT_DIR/config.yml -logpath=$SCRIPT_DIR/logs > /dev/null 2>&1 &
  for ((i=0; i<5; ++i))
    do
      sleep 1
      if ! pgrep $EXECUTABLE_FILE > /dev/null; then
        echo -e ".\c"
      else
        echo 'Start success!'
        break;
      fi
    done
  if ! pgrep $EXECUTABLE_FILE > /dev/null; then
    echo "Start fail!"
  fi
}

case $1 in
    "ps" )
        Ps
    ;;
    "start" )
        Start
    ;;
    "stop" )
       Stop
    ;;
    "restart" )
       Stop
       Start
    ;;
    * )
        echo "unknown command"
        exit 1
    ;;
esac