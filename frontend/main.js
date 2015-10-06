require.config({
    paths: {
        "JSXTransformer": "components/react/JSXTransformer",
        "immutable": "components/immutable/dist/immutable",
        "jquery": "components/jquery/dist/jquery",
        "jsx": "components/requirejs-react-jsx/jsx",
        "morearty": "components/moreartyjs/dist/morearty",
        "react": "components/react/react-with-addons",
        "socket.io": "components/socket.io-client/socket.io",
        "text": "components/requirejs-text/text",
        "moment": "components/moment/moment",
    },

    shim : {
        "react": {
            "exports": "React"
        },
        "JSXTransformer": "JSXTransformer"
    },

    jsx: {
        fileExtension: ".jsx",
        transformOptions: {
            harmony: true,
            stripTypes: false
        },
        usePragma: false
    }
});

require(['react', 'immutable'], function (React, Imm) {
    window.React = React;
    window.Immutable = Imm;

    require(['app/Bootstrap'], function (Bootstrap) {
        React.render(
            React.createFactory(Bootstrap)(),
            document.getElementById('root')
        );
    });
});
