# Enable Powerlevel10k instant prompt. Should stay close to the top of ~/.zshrc.
# Initialization code that may require console input (password prompts, [y/n]
# confirmations, etc.) must go above this block; everything else may go below.
if [[ -r "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh" ]]; then
  source "${XDG_CACHE_HOME:-$HOME/.cache}/p10k-instant-prompt-${(%):-%n}.zsh"
fi

# If you come from bash you might have to change your $PATH.
export GOPATH=/Users/ribincao/Go
export PATH=$HOME/bin:/usr/local/bin:$PATH:$GOPATH/bin
export CARGO_NET_GIT_FETCH_WITH_CLI=true

# Path to your oh-my-zsh installation.
export ZSH="$HOME/.oh-my-zsh"

# Set name of the theme to load --- if set to "random", it will
# load a random theme each time oh-my-zsh is loaded, in which case,
# to know which specific one was loaded, run: echo $RANDOM_THEME
# See https://github.com/ohmyzsh/ohmyzsh/wiki/Themes
# ZSH_THEME="robbyrussell"
ZSH_THEME="powerlevel10k/powerlevel10k"
POWERLEVEL9K_MODE="awesome-patched"

# Set list of themes to pick from when loading at random
# Setting this variable when ZSH_THEME=random will cause zsh to load
# a theme from this variable instead of looking in $ZSH/themes/
# If set to an empty array, this variable will have no effect.
# ZSH_THEME_RANDOM_CANDIDATES=( "robbyrussell#"agnoster#)

# Uncomment the following line to use case-sensitive completion.
# CASE_SENSITIVE="true"

# Uncomment the following line to use hyphen-insensitive completion.
# Case-sensitive completion must be off. _ and - will be interchangeable.
# HYPHEN_INSENSITIVE="true"

# Uncomment one of the following lines to change the auto-update behavior
# zstyle ':omz:update' mode disabled  # disable automatic updates
# zstyle ':omz:update' mode auto      # update automatically without asking
# zstyle ':omz:update' mode reminder  # just remind me to update when it's time

# Uncomment the following line to change how often to auto-update (in days).
# zstyle ':omz:update' frequency 13

# Uncomment the following line if pasting URLs and other text is messed up.
# DISABLE_MAGIC_FUNCTIONS="true"

# Uncomment the following line to disable colors in ls.
# DISABLE_LS_COLORS="true"

# Uncomment the following line to disable auto-setting terminal title.
# DISABLE_AUTO_TITLE="true"

# Uncomment the following line to enable command auto-correction.
# ENABLE_CORRECTION="true"

# Uncomment the following line to display red dots whilst waiting for completion.
# You can also set it to another string to have that shown instead of the default red dots.
# e.g. COMPLETION_WAITING_DOTS="%F{yellow}waiting...%f"
# Caution: this setting can cause issues with multiline prompts in zsh < 5.7.1 (see #5765)
# COMPLETION_WAITING_DOTS="true"

# Uncomment the following line if you want to disable marking untracked files
# under VCS as dirty. This makes repository status check for large repositories
# much, much faster.
# DISABLE_UNTRACKED_FILES_DIRTY="true"

# Uncomment the following line if you want to change the command execution time
# stamp shown in the history command output.
# You can set one of the optional three formats:
# "mm/dd/yyyy"|"dd.mm.yyyy"|"yyyy-mm-dd"
# or set a custom format using the strftime function format specifications,
# see 'man strftime' for details.
# HIST_STAMPS="mm/dd/yyyy"

# Would you like to use another custom folder than $ZSH/custom?
# ZSH_CUSTOM=/path/to/new-custom-folder

# Which plugins would you like to load?
# Standard plugins can be found in $ZSH/plugins/
# Custom plugins may be added to $ZSH_CUSTOM/plugins/
# Example format: plugins=(rails git textmate ruby lighthouse)
# Add wisely, as too many plugins slow down shell startup.
plugins=(
    git
    zsh-syntax-highlighting
    zsh-autosuggestions
    vscode
)

source $ZSH/oh-my-zsh.sh

# User configuration

# export MANPATH="/usr/local/man:$MANPATH"

# You may need to manually set your language environment
# export LANG=en_US.UTF-8

# Preferred editor for local and remote sessions
# if [[ -n $SSH_CONNECTION ]]; then
#   export EDITOR='vim'
# else
#   export EDITOR='mvim'
# fi

# Compilation flags
# export ARCHFLAGS="-arch x86_64"

# Set personal aliases, overriding those provided by oh-my-zsh libs,
# plugins, and themes. Aliases can be placed here, though oh-my-zsh
# users are encouraged to define aliases within the ZSH_CUSTOM folder.
# For a full list of active aliases, run `alias`.
#
# Example aliases
# alias zshconfig="mate ~/.zshrc"
# alias ohmyzsh="mate ~/.oh-my-zsh"
alias cd-root="cd /Users/ribincao"
alias cd-desk="cd /Users/ribincao/Desktop"
alias cd-ribin="cd /Users/ribincao/Desktop/ribin-workspace"
alias cd-bud="cd /Users/ribincao/Desktop/bud-workspace"
alias cd-tool="cd /Users/ribincao/Desktop/bud-workspace/bud-tools"
alias cd-china="cd /Users/ribincao/Desktop/bud-workspace/bud-engine-china"
alias cd-us="cd /Users/ribincao/Desktop/bud-workspace/bud-engine-us"
alias cd-common="cd /Users/ribincao/Desktop/bud-workspace/bud-engine-common"

alias cd-log="cd /Users/ribincao/Desktop/bud-workspace/bud-logs"
alias cd-refactor="cd /Users/ribincao/Desktop/bud-workspace/bud-engine-refactor"
alias cd-scripts="cd /Users/ribincao/Desktop/bud-workspace/bud-tools/scripts"
alias python="/opt/homebrew/bin/python3"

alias context_change="/Users/ribincao/Desktop/bud-workspace/bud-tools/scripts/change_context.sh"
alias build_fleet="/Users/ribincao/Desktop/bud-workspace/bud-tools/scripts/build_fleet.sh"
alias build_fleet_cn="/Users/ribincao/Desktop/bud-workspace/bud-tools/scripts/build_fleet_china.sh"
alias dumplog="/Users/ribincao/Desktop/bud-workspace/bud-tools/scripts/dumplog.sh"
alias checkout="/Users/ribincao/Desktop/bud-workspace/bud-tools/scripts/checkout.sh"
alias apollo="/Users/ribincao/Desktop/bud-workspace/bud-tools/scripts/apollo.sh"
alias engine_tool="cd /Users/ribincao/Desktop/bud-workspace/bud-tools/scripts && ./tool.sh"

alias current_context="kubectl config current-context"
alias search="grep --color -n -E -re"
alias vim="/opt/homebrew/bin/nvim"
alias vim-zsh="vim ~/.zshrc"
alias k8s="/opt/homebrew/bin/kubectl"
alias help="/opt/homebrew/bin/tldr"

alias du="/opt/homebrew/bin/dust"
alias df="/opt/homebrew/bin/duf"
alias ls="/opt/homebrew/bin/exa"
alias grep-color="/opt/homebrew/bin/rg"
alias cat="/opt/homebrew/bin/bat"
alias ts="date '+%s'"

alias build_test="cd /Users/$(whoami)/Desktop/bud-refactor/engine-server-logic-test && bash build.sh v1.80.0-master us-west-1 /Users/$(whoami)/Desktop/bud-workspace/BUD_GAME_GLOBAL m1"
source ~/.oh-my-zsh/custom/plugins/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh

eval "$(thefuck --alias)"
eval "$(zoxide init zsh)"
export HOMEBREW_GITHUB_API_TOKEN=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

# To customize prompt, run `p10k configure` or edit ~/.p10k.zsh.
[[ ! -f ~/.p10k.zsh ]] || source ~/.p10k.zsh

# python path
PYTHONPATH="/Users/ribincao/Desktop/ribin-workspace/ribin-py-2dgame:/Users/ribincao/Desktop/ribin-workspace/ribin-chatGPT"
export PYTHONPATH
