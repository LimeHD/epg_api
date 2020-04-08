lock '3.12.1'

set :application, 'epg_api'
set :user, 'master'

set :repo_url, 'git@github.com:LimeHD/epg_api.git' if ENV['USE_LOCAL_REPO'].nil?

# Тут будут линки на конфиг
# set :linked_files, %w(app/includes/config.inc.php config/database.yml app/includes/config_newtickets.inc.php app/includes/config_rate.inc.php app/includes/config_sleep.yml)
set :linked_dirs, %w(log)

if ENV['BRANCH'].nil?
  ask :branch, proc { `git rev-parse --abbrev-ref HEAD`.chomp }
else
  set :branch, ENV['BRANCH']
end

set :deploy_to, -> { "/home/#{fetch(:user)}/#{fetch(:application)}" }

namespace :deploy do
  # after 'updated', :build_go
end

GO_BIN='/opt/go/1.13.9/bin/go'
desc 'Build script on golang'
task :build_go do
  on release_roles(:app) do
    execute "cd #{release_path}; #{GO_BIN} get && #{GO_BIN} build"
  end
end
