# Repository
mockery --name=DBRepository --dir ./repository --output ./mocks/repository
mockery --name=RedisRepository --dir ./repository --output ./mocks/repository

# Service
mockery --name=Service --dir ./service --output ./mocks/service