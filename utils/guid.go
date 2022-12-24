package utils

import (
  "crypto/md5"
  "crypto/rand"
  "encoding/binary"
  "encoding/hex"
  "fmt"
  "io"
  "os"
  "sync/atomic"
  "time"
)

var objectIdCounter uint32 = 0

var machineId = readMachineId()

type ObjectId string

func readMachineId() []byte {
  var sum [3]byte
  id := sum[:]
  hostname, err1 := os.Hostname()
  if err1 != nil {
    _, err2 := io.ReadFull(rand.Reader, id)
    if err2 != nil {
      panic(fmt.Errorf("cannot get hostname: %v; %v", err1, err2))
    }
    return id
  }
  hw := md5.New()
  hw.Write([]byte(hostname))
  copy(id, hw.Sum(nil))
  fmt.Println("readMachineId:" + string(id))
  return id
}

// GUID returns a new unique ObjectId.
// 4byte 时间，
// 3byte 机器ID
// 2byte pid
// 3byte 自增ID
func GetGUID() ObjectId {
  var b [12]byte
  // Timestamp, 4 bytes, big endian
  binary.BigEndian.PutUint32(b[:], uint32(time.Now().Unix()))
  // Machine, first 3 bytes of md5(hostname)
  b[4] = machineId[0]
  b[5] = machineId[1]
  b[6] = machineId[2]
  // Pid, 2 bytes, specs don't specify endianness, but we use big endian.
  pid := os.Getpid()
  b[7] = byte(pid >> 8)
  b[8] = byte(pid)
  // Increment, 3 bytes, big endian
  i := atomic.AddUint32(&objectIdCounter, 1)
  b[9] = byte(i >> 16)
  b[10] = byte(i >> 8)
  b[11] = byte(i)
  return ObjectId(b[:])
}

// Hex returns a hex representation of the ObjectId.
// 返回16进制对应的字符串
func (id ObjectId) Hex() string {
  return hex.EncodeToString([]byte(id))
}
