define(['react', 'morearty', 'jquery'], function(React, Morearty, $) {
    return React.createClass({
        displayName: 'PingView',
        mixins: [Morearty.Mixin],
        onChange: function(e) {
            return;
            var editBox = e.target;
            var binding = this.getDefaultBinding();
            var selected = binding.get('selected_request');
            if (selected === undefined) {
                return;
            }
            var request = binding.sub('history').sub(selected);
            request.sub(this.props.view).update(function(text) {
                return editBox.innerText;
            });
        },
        render: function() {
            var binding = this.getDefaultBinding();
            var request = binding.sub('history').sub(binding.get('selected_request'));
            var text = request.get(this.props.view);
            return (
                <div className="column column-flex">
                </div>
            );
        },
    });
});
