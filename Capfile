require 'rubygems'
require 'bundler'
Bundler.setup(:deploy)

# Load DSL and Setup Up Stages
require 'capistrano/setup'

# Includes default deployment tasks
require 'capistrano/deploy'

require "capistrano/scm/git"
install_plugin Capistrano::SCM::Git

require 'capistrano/shell'

require 'capistrano/systemd/multiservice'
install_plugin Capistrano::Systemd::MultiService.new_service('go', service_type: 'user')
