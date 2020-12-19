#!/bin/zsh
scriptDir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
catalina_cc_path="/usr/local/bin/gcc-9"
catalina_cxx_path="/usr/local/bin/g++-9"

function main() {
  if [[ -d $scriptDir/build ]]; then
    echo "WARN: Deleting $scriptDir/build directory..."
    rm -rf $scriptDir/build
  fi
  if command -v pip3 >/dev/null 2>&1 ; then
    if [[ ! -f ~/Library/Python/3.7/bin/conan ]]; then
      echo "INSTALL conan package manager ..."
      pip3 install --user conan #conan will be added in ~/Library/Python/3.7/bin/conan. Needs new shell
    fi 
  else
    echo "WARN: can't find pip3!"
    exit -1
  fi
  echo "INSTALL latest gcc ..."
  brew install gcc || brew upgrade gcc
  if [[ $isMacos ]]; then
    [[ -f $catalina_cc_path ]] && export CC=$catalina_cc_path || error_exit "Cannot find GNU compiler at $catalina_c_path" 1
    [[ -f $catalina_cxx_path ]] && export CXX=$catalina_cxx_path || error_exit "Cannot find GNU compiler at $catalina_cxx_path" 1
  fi
  [[ ! -d $scriptDir/build ]] && mkdir -p $scriptDir/build
  if [[ -f ~/Library/Python/3.7/bin/conan ]]; then
    cd build
    [[ -f ~/conan/profiles/default ]] && rm ~/.conan/profiles/default
    conan profile new default --detect  # Generates default profile detecting GCC and sets old ABI
    conan profile update settings.compiler.libcxx=libstdc++11 default  # Sets libcxx to C++11 ABI
    conan install ..
  else
    error_exit "Conan not installed!" 2
  fi
  cmake -H$scriptDir -B $scriptDir/build  -DCMAKE_CXX_COMPILER=$CXX 
}

function detect_os() {
  uname=`uname`
  if [[ $uname == 'Darwin' ]]; then
    export isMacos=true
  elif [[ $uname == 'Linux' ]]; then
    export isLinux=true
  elif [[ $uname == CYGWIN* ]]; then
    export isCygwin=true
  else
    echo "This Setup Script only works for Windows Cygwin, Windows Subsystem for Linux and MacOS"
    exit -1
  fi
  export osreleaseFile=/proc/sys/kernel/osrelease
  if [[ -f  $osreleaseFile ]]; then
    osrelease=`cat $osreleaseFile`
    if [[ $osrelease == *Microsoft* ]]; then
      export isWsl=true
    fi
  fi
}

detect_os
main
