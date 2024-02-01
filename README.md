## i2c-scanner
I2C bus scanner for linux.
Written in pure Go, can be compiled to a static binary without dependencies.

### installation
```
go install github.com/damdo/i2c-scanner@latest
```

### usage
```
i2c-scanner
```

the output will look similar to this:
```
$ i2c-scanner

I2C buses available:
- name: /dev/i2c-0
  number: 0
  aliases: I2C0
  pins:
   * SDA: I2C0_SDA(GPIO0)
   * SCL: I2C0_SCL(GPIO1)

- name: /dev/i2c-1
  number: 1
  aliases: I2C1
  pins:
   * SDA: I2C1_SDA(GPIO2)
   * SCL: I2C1_SCL(GPIO3)

- name: /dev/i2c-2
  number: 2
  aliases: I2C2
  pins:
   * SDA: INVALID
   * SCL: INVALID

- name: /dev/i2c-3
  number: 3
  aliases: I2C3
  pins:
   * SDA: INVALID
   * SCL: INVALID
```

### credits

Code heavily inspired by: https://pkg.go.dev/periph.io/x/periph/conn/i2c
