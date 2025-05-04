#!/usr/bin/env bash

host="$1"
shift
cmd="$@"

echo "⏳ Waiting for $host to be ready..."

until nc -z $host 2>/dev/null; do
  echo "🚧 Still waiting for $host..."
  sleep 1
done

echo "✅ $host is up. Starting application..."
sleep 1  

exec $cmd#!/usr/bin/env bash

host="$1"
shift
cmd="$@"

until nc -z ${host} 2>/dev/null; do
  echo "Waiting for ${host}..."
  sleep 1
done

exec $cmd
