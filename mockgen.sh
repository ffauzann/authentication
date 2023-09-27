# Repository
mockery --disable-version-string --name=DBRepository --dir ./internal/repository --output ./mocks/internal/repository
mockery --disable-version-string --name=RedisRepository --dir ./internal/repository --output ./mocks/internal/repository

# Service
mockery --disable-version-string --name=Service --dir ./internal/service --output ./mocks/internal/service