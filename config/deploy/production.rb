set :stage, :production
set :disallow_pushing, true
server '194.35.48.30', user: fetch(:user), port: '22', roles: %w(app).freeze
set :keep_releases, 10
