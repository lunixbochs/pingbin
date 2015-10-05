define(['react', 'morearty', 'jsx!app/PingListItem'], function(React, Morearty, PingListItem) {
    return React.createClass({
        displayName: 'PingList',
        mixins: [Morearty.Mixin],
        componentDidMount: function() {
            this.handleResize();
            window.addEventListener('resize', this.handleResize);
        },
        componentWillUnmount: function() {
            window.removeEventListener('resize', this.handleResize);
        },
        handleResize: function() {
            var binding = this.getDefaultBinding();
            var w = binding.sub('window');
            w.set('width', window.innerWidth);
            w.set('height', window.innerHeight);
        },
        shouldComponentUpdateOverride: function(orig, nextProps, nextState) {
            var binding = this.getDefaultBinding();
            if (binding.get('history_pos') !== nextProps.binding.get('history_pos')) {
                return true;
            }
            return orig();
        },
        onScroll: function(e) {
            var itemHeight = 41;
            var binding = this.getDefaultBinding();
            var pos = Math.floor(e.target.scrollTop / itemHeight);
            binding.set('history_pos', pos);
        },
        render: function() {
            var itemHeight = 41;
            var binding = this.getDefaultBinding();
            var historyBinding = binding.sub('history');
            var history = historyBinding.get();
            var first = 0;
            var last = history.count();
            var pushTop = 0;
            var pushBot = 0;
            var visible = history;
            var renderPing = function(item, index) {
                index += first;
                var itemBinding = historyBinding.sub(index);
                var onClick = function() {
                    binding.set('selected_request', index);
                }
                var selected = index === binding.get('selected_request');
                return <PingListItem
                        onClick={onClick} selected={selected} binding={itemBinding} key={itemBinding.toJS('id')} />
            }
            if (history.count() > 0) {
                var section = 40;
                var pos = (binding.get('history_pos') || 0);
                pos = Math.floor(pos / 2) * 2;
                first = Math.max(pos - section / 2, 0);
                last = Math.min(pos + section / 2, history.count());
                pushTop = first * itemHeight + "px";
                pushBot = (history.count() - last) * itemHeight + "px";
                visible = history.slice(first, last);
            }
            return (
                <div className="list column column-narrow" onScroll={this.onScroll} tabIndex="0">
                    <div className="push" style={{height: pushTop}}></div>
                    {visible.map(renderPing).toArray()}
                    <div className="push" style={{height: pushBot}}></div>
                </div>
            );
        },
    });
});
