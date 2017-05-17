// This module exports functions that give access to the goa simple sample API hosted at localhost:8080.
// It uses the axios javascript library for making the actual HTTP requests.
define(['axios'] , function (axios) {
  function merge(obj1, obj2) {
    var obj3 = {};
    for (var attrname in obj1) { obj3[attrname] = obj1[attrname]; }
    for (var attrname in obj2) { obj3[attrname] = obj2[attrname]; }
    return obj3;
  }

  return function (scheme, host, timeout) {
    scheme = scheme || 'http';
    host = host || 'localhost:8080';
    timeout = timeout || 20000;

    // Client is the object returned by this module.
    var client = axios;

    // URL prefix for all API requests.
    var urlPrefix = scheme + '://' + host;

  // 追加
  // path is the request path, the format is "/api/v1/accounts"
  // email, name are used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.addAccounts = function (path, email, name, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
      params: {
        email: email,
        name: name
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 追加
  // path is the request path, the format is "/api/v1/accounts_data"
  // email, name are used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.addAccounts_data = function (path, email, name, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
      params: {
        email: email,
        name: name
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 追加
  // path is the request path, the format is "/api/v1/bottles"
  // account_id, name, quantity are used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.addBottles = function (path, account_id, name, quantity, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
      params: {
        account_id: account_id,
        name: name,
        quantity: quantity
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 追加
  // path is the request path, the format is "/api/v1/bottles_data"
  // account_id, name, quantity are used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.addBottles_data = function (path, account_id, name, quantity, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'post',
      params: {
        account_id: account_id,
        name: name,
        quantity: quantity
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // ユーザー（配列）
  // path is the request path, the format is "/api/v1/response/users/array"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.arrayResponse = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 削除
  // path is the request path, the format is "/api/v1/accounts/:id"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.deleteAccounts = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'delete',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 削除
  // path is the request path, the format is "/api/v1/accounts_data/:id"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.deleteAccounts_data = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'delete',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 削除
  // path is the request path, the format is "/api/v1/bottles/:id"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.deleteBottles = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'delete',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 削除
  // path is the request path, the format is "/api/v1/bottles_data/:id"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.deleteBottles_data = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'delete',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // ちょっと特殊ケース
  // path is the request path, the format is "/api/v1/method/users/:id/follow/:type"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.etcMethod = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // フォロー操作
  // path is the request path, the format is "/api/v1/method/users/follow"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.followMethod = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'put',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // ユーザー（ハッシュ）
  // path is the request path, the format is "/api/v1/response/users/hash"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.hashResponse = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 挨拶する
  // path is the request path, the format is "/api/v1/actions/hello"
  // name is used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.helloActions = function (path, name, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      params: {
        name: name
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 複数アクション（:id）
  // path is the request path, the format is "/api/v1/actions/:id"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.idActions = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 複数
  // path is the request path, the format is "/api/v1/accounts"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.listAccounts = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 複数
  // path is the request path, the format is "/api/v1/accounts_data"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.listAccounts_data = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 複数
  // path is the request path, the format is "/api/v1/bottles"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.listBottles = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 複数
  // path is the request path, the format is "/api/v1/bottles_data"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.listBottles_data = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // リストを返す
  // path is the request path, the format is "/api/v1/method/list"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.listMethod = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // ユーザー（複数）
  // path is the request path, the format is "/api/v1/response/users"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.listResponse = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 複数(リレーション版)
  // path is the request path, the format is "/api/v1/bottles/relation"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.listRelationBottles = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // HTTPメソッド
  // path is the request path, the format is "/api/v1/method/get"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.methodMethod = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // ネストしたMediaType
  // path is the request path, the format is "/api/v1/response/users/nested"
  // test is used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.nestedResponse = function (path, test, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      params: {
        test: test
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // サーバーとの導通確認
  // path is the request path, the format is "/api/v1/actions/ping"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.pingActions = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // セキュリティの例です
  // path is the request path, the format is "/api/v1/securiy"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.securitySecurity = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 単数
  // path is the request path, the format is "/api/v1/accounts/:id"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.showAccounts = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 単数
  // path is the request path, the format is "/api/v1/accounts_data/:id"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.showAccounts_data = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 単数
  // path is the request path, the format is "/api/v1/bottles/:id"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.showBottles = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 単数
  // path is the request path, the format is "/api/v1/bottles_data/:id"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.showBottles_data = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // ユーザー（単数）
  // path is the request path, the format is "/api/v1/response/users/:id"
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.showResponse = function (path, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 更新
  // path is the request path, the format is "/api/v1/accounts/:id"
  // email, name are used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.updateAccounts = function (path, email, name, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'put',
      params: {
        email: email,
        name: name
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 更新
  // path is the request path, the format is "/api/v1/accounts_data/:id"
  // email, name are used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.updateAccounts_data = function (path, email, name, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'put',
      params: {
        email: email,
        name: name
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 更新
  // path is the request path, the format is "/api/v1/bottles/:id"
  // name, quantity are used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.updateBottles = function (path, name, quantity, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'put',
      params: {
        name: name,
        quantity: quantity
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // 更新
  // path is the request path, the format is "/api/v1/bottles_data/:id"
  // name, quantity are used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.updateBottles_data = function (path, name, quantity, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'put',
      params: {
        name: name,
        quantity: quantity
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }

  // Validation
  // path is the request path, the format is "/api/v1/validation"
  // defaultType, email, enumType, id, integerType, reg, stringType are used to build the request query string.
  // config is an optional object to be merged into the config built by the function prior to making the request.
  // The content of the config object is described here: https://github.com/mzabriskie/axios#request-api
  // This function returns a promise which raises an error if the HTTP response is a 4xx or 5xx.
  client.validationValidation = function (path, defaultType, email, enumType, id, integerType, reg, stringType, config) {
    cfg = {
      timeout: timeout,
      url: urlPrefix + path,
      method: 'get',
      params: {
        defaultType: defaultType,
        email: email,
        enumType: enumType,
        id: id,
        integerType: integerType,
        reg: reg,
        stringType: stringType
      },
      responseType: 'json'
    };
    if (config) {
      cfg = merge(cfg, config);
    }
    return client(cfg);
  }
  return client;
  };
});
