set :stage, :reproduction
set :disallow_pushing, true
server ENV['REPRODUCTION_HOST'], user: fetch(:user), port: '22', roles: %w(app).freeze
set :keep_releases, 10
