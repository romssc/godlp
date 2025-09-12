#!/usr/bin/env bash

# ========== FLAGS ============

set -e

# ========== IMPORTNANT VARIABLES ============

VENV_DIR="$PKG_DIR/.py/venv"
OS="$(uname -s)"

# ========== FUNCTIONS ============

fail() {
    echo "$1"
    exit 1
}

run_step() {
    err=$(mktemp)
    "$@" 2>"$err" || fail "failed check: $(<"$err")"
    rm -f "$err"
}

install_python_mac() {
    command -v brew >/dev/null 2>&1 || /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)" || fail "failed to install python: not able to install homebrew"
    brew install python || fail "failed to install python: homebrew ran into a problem"
}

install_python_linux() {
    if command -v apt >/dev/null 2>&1; then
        sudo apt update || fail "failed to install python: not able to update apt"
        sudo apt install -y python3 python3-venv python3-pip || fail "failed to install python: apt ran into a problem"
    else
        fail "failed to install python: linux distro is not supported"
    fi
}

install_python_windows() {
    command -v winget >/dev/null 2>&1 || fail "failed to install python: winget not found"
    winget install --id Python.Python.3 -e --source winget || fail "failed to install python: winget ran into a problem"
}

# ========== STEPS ============

# 1. Check if Python3 is available
run_step command -v python3 >/dev/null 2>&1

# 2. Create venv if not exists
[ -d "$VENV_DIR" ] || run_step python3 -m venv "$VENV_DIR"

# 3. Activate venv
source "$VENV_DIR/bin/activate" || fail "failed check: not able to activate venv"

PYTHON="$VENV_DIR/bin/python"
PIP="$VENV_DIR/bin/pip"

# 4. Upgrade pip
run_step "$PYTHON" -m pip install --upgrade pip

# 6. Install requirements
run_step "$PIP" install --upgrade yt-dlp

# 7. Test Python & yt-dlp in venv
run_step "$PYTHON" -m yt_dlp --version

exit 0