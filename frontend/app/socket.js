define(['socket.io'], function (io) {
    var url = window.location.href.replace(/[^\/]+$/, '');
    var s = io(url);
    window.addEventListener('unload', function() {s.disconnect()});
    s.on('connect', function() {
        console.log('socket connected');
    });
    s.on('disconnection', function() {
        console.log('socket disconnected');
    });
    // TODO: connection error? UI status? reconnecting?
    return s;
});
