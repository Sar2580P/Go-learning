* Command for checking race condition

    ```go test ./... -v -count=1 -race```

---

## BookingStore Implementations Comparison

| Feature | MemoryStore | ConcurrentStore | RedisStore |
|---------|-------------|-----------------|-----------|
| **Storage** | Map (in-memory) | Map (in-memory) | Redis (distributed) |
| **Thread-Safe** | ❌ No | ✅ Yes (RWMutex) | ✅ Yes (atomic ops) |
| **Persistence** | ❌ Lost on restart | ❌ Lost on restart | ✅ Persistent |
| **TTL Support** | ❌ Manual cleanup | ❌ Manual cleanup | ✅ Auto-expire |
| **Scalability** | Single instance | Single instance | Distributed |
| **Hold + Confirm** | Basic map storage | Mutex-protected map | 2-key architecture |
| **Use Case** | Learning/testing | Dev/single-instance | Production |

**MemoryStore**: Simple map, no synchronization → race conditions with concurrent access  
**ConcurrentStore**: RWMutex protects map → safe for concurrent reads/writes  
**RedisStore**: Distributed with auto-TTL and reverse lookup → production-grade with cleanup  

---

# Redis Two-Hierarchy Key Pattern

## Key Structure
```
seat:{movieID}:{seatID}  → Booking JSON (with/without TTL)
session:{sessionID}      → "seat:{movieID}:{seatID}" (reverse pointer)
```

## Why Two Keys?

**Forward lookup** (`seat:movieID:seatID`): Check availability, list bookings  
**Reverse lookup** (`session:sessionID`): Find which seat a session owns

Without reverse lookup, expired held seats can't be cleaned up—you'd have to scan all seats.

## Flow

1. **Hold**: Create both keys with same TTL (2 min)
2. **Confirm**: Remove TTL via PERSIST on both keys → permanent booking
3. **Expire**: Both keys auto-delete when TTL runs out → seat available again

## Code Pattern
```go
sessionID := uuid.New().String()
seatKey := fmt.Sprintf("seat:%s:%s", movieID, seatID)

// Forward + Reverse, same TTL
rdb.SetArgs(ctx, seatKey, bookingJSON, redis.SetArgs{TTL: 2*time.Minute})
rdb.Set(ctx, "session:"+sessionID, seatKey, 2*time.Minute)

// Confirm: remove TTL from both
rdb.Persist(ctx, seatKey)
rdb.Persist(ctx, "session:"+sessionID)
```

Also see [`getSession`](internal/booking/redis_store.go#L141) in redis_store.go

## Use Cases
- Temporary reservations with auto-expiration (seats, carts, tickets)
- Session state that needs bidirectional lookup
- Auto-cleanup without external jobs
