# superplace


日志配置格式
```log
{
  "logger": {
    "game_log": {
      "level": "debug",
      "stack_level": "error",
      "enable_console": true,
      "enable_write_file": false,
      "max_age": 7,
      "time_format": "15:04:05.000",
      "print_caller": true,
      "rotation_time": 3600,
      "@rotation_time": [
        "3600(秒):每小时分割日志",
        "86400(秒):每天分割日志"
      ],
      "file_link_path": "logs/game.link",
      "file_path_format": "logs/game_%Y%m%d%H%M.log"
    },
    "master_log": {
      "level": "debug",
      "enable_console": true,
      "enable_write_file": false,
      "file_link_path": "logs/master.log",
      "file_path_format": "logs/master_%Y%m%d%H%M.log"
    }
    "test_handler": {
      "level": "debug",
      "enable_console": false,
      "enable_write_file": true,
      "file_link_path": "logs/test_handler.log",
      "file_path_format": "logs/test_handler_%Y%m%d%H%M.log"
    }
  }
}
```


服务配置格式
```master center game web
{
  "env": "gc",
  "debug": true,
  "print_level": "info",
   "include": [
    "common/logger.json",
    "profile-dev-cluster.json",
    "profile-dev-data-config.json"
  ],
  "cluster": {
    "discovery": {
      "mode": "nats",
      "@mode1": "mode=default,从profile-{x}.json读取node节点的配置数据",
      "@mode2": "mode=nats,通过nats->master_node_id获取已注册的节点",
      "@mode3": "mode=etcd,通过etcd同步已注册节点"
    },
    "nats": {
      "master_node_id": "gc-master",
      "address": "nats://127.0.0.1:4222",
      "reconnect_delay": 1,
      "max_reconnects": 0,
      "request_timeout": 2,
      "user": "",
      "password": ""
    }
  },
  "node": {
    "master": [
      {
        "node_id": "gc-master",
        "address": "",
        "__settings__": {
          "ref_logger": "master_log"
        },
        "enable": true
      }
    ],
    "center": [
      {
        "node_id": "gc-center",
        "address": "",
        "__settings__": {
          "ref_logger": "center_log"
        },
        "enable": true
      }
    ],
    "gate": [
      {
        "node_id": "gc-gate-1",
        "address": ":10010",
        "__settings__": {
          "db_id_list" : {
            "game_db_id": "game_db_1"
          },
          "ref_logger": "gate_log"
        },
        "enable": true
      }
    ],
    "web": [
      {
        "node_id": "gc-web-1",
        "address": "127.0.0.1:8089",
        "__settings__": {
          "ref_logger": "master_log"
        },
        "enable": true
      }
    ],
    "game": [
      {
        "node_id": "10001",
        "__settings__": {
          "db_id_list" : {
            "game_db_id": "game_db_1"
          },
          "ref_logger": "game_log"
        },
        "enable": true
      }
    ]
  },
  "logger": {
    "master_log": {
      "level": "debug",
      "enable_console": true,
      "enable_write_file": false,
      "file_link_path": "logs/master.log",
      "file_path_format": "logs/master_%Y%m%d%H%M.log"
    },
    "center_log": {
      "level": "debug",
      "enable_console": true,
      "enable_write_file": false,
      "file_link_path": "logs/center.log",
      "file_path_format": "logs/center_%Y%m%d%H%M.log"
    },
    "gate_log": {
      "level": "debug",
      "enable_console": true,
      "enable_write_file": false,
      "file_link_path": "logs/gate.log",
      "file_path_format": "logs/gate_%Y%m%d%H%M.log"
    },
    "game_log": {
      "level": "debug",
      "enable_console": true,
      "enable_write_file": false,
      "file_link_path": "logs/game.log",
      "file_path_format": "logs/game_%Y%m%d%H%M.log"
    },
    "cross_log": {
      "level": "debug",
      "enable_console": true,
      "enable_write_file": false,
      "file_link_path": "logs/cross.log",
      "file_path_format": "logs/cross_%Y%m%d%H%M.log"
    }
  },
  "data_config": {
    "parser": "json",
    "data_source": "file",
    "file": {
      "file_path": "data/",
      "ext_name": ".json",
      "reload_time": 3000
    },
    "redis": {
      "prefix_key": "data_config",
      "subscribe_key": "data_config_change",
      "address": "127.0.0.1:6379",
      "password": "",
      "db": 7
    }
  }
}
```