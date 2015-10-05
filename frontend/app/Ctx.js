define(['morearty', 'immutable'], function(Morearty, Immutable) {
    return Morearty.createContext({
        initialState: {
            'history': [
                {
                    'type': 'ICMP',
                    'ip': '127.0.0.1',
                    'time': 0,
                    'token': '11223344',
                    'packet': 'packet',
                },
                {
                    'type': 'DNS',
                    'ip': '127.0.0.1',
                    'time': 0,
                    'token': '11223344',
                    'packet': 'packet',
                    'domain': 'localhost.',
                },
                {
                    'type': 'HTTP',
                    'ip': '127.0.0.1',
                    'time': 0,
                    'token': '11223344',
                    'packet': 'packet',
                    'domain': 'localhost',
                    'path': '/index.html',
                    'headers': {
                        'User-Agent': 'Google Blarg (webkit-compatible)',
                    },
                },
            ],
            'window': {'width': 1, 'height': 1},
        },
    });
});
