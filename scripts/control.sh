#!/bin/bash

# 获取操作系统类型
OS=$(uname)
echo "当前操作系统为${OS}"

scriptFile=""
if [ "${OS}" == "Darwin" ]; then
    echo "Running on macOS"
    scriptFile="mac_main"
elif [ "${OS}" == "Linux" ]; then
    echo "Running on Linux"
    scriptFile="linux_main"
else
    echo "Unsupported operating system: $OS"
    exit 1
fi

Ps(){
  pid=$(pgrep $scriptFile)
  if [ -n "$pid" ] ; then
    echo "进程已启动: $pid"
  else
    echo "进程未启动"
  fi
}

Stop(){
  # 检查是否存在进程
  if pgrep $scriptFile > /dev/null; then
      # 如果存在则通过kill命令杀死进程
      echo "开始关闭"
      pkill $scriptFile
      for ((i=0; i<10; ++i))
        do
          sleep 1
          if pgrep $scriptFile > /dev/null; then
            echo -e ".\c"
          else
            echo 'Stop Success!'
            break;
          fi
        done

      if pgrep $scriptFile > /dev/null; then
        echo 'Kill Process!'
        pkill -9 $scriptFile
      fi
  else
      echo "进程已关闭"
  fi
}

Start(){
  echo "nohup ./$scriptFile > /dev/null 2>&1 &"
  nohup ./$scriptFile -config=config.yml > /dev/null 2>&1 &
  for ((i=0; i<5; ++i))
    do
      sleep 1
      if ! pgrep $scriptFile > /dev/null; then
        echo -e ".\c"
      else
        echo 'Start success!'
        break;
      fi
    done
  if ! pgrep $scriptFile > /dev/null; then
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