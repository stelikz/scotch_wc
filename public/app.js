new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
      	username: null, // Our username
        joined: false // True if email and username have been filled in
    },

    created: function() {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        // if statement to iterate if there is 
        this.ws.addEventListener('message', function(e) {
            var msg = JSON.parse(e.data);
            for (i = 0; i < msg.length; i++) {
                self.chatContent += '<div class="chip">'
                    + msg[i].username
                + '</div>'
                + msg[i].message + '<br/>';

            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight;
                }
            self.chatContent += '<div class="chip">'
                    + msg.username
                + '</div>'
                + msg.message + '<br/>'; // Parse emojis

            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
        });
    },

    methods: {
        send: function () {
            if (this.newMsg != '') {
                this.ws.send(
                    JSON.stringify({
                        username: this.username,
                        message: $('<p>').html(this.newMsg).text() // Strip out html
                    }
                ));
                this.newMsg = ''; // Reset newMsg
            }
        },

        join: function () {
            if (!this.username) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            this.username = $('<p>').html(this.username).text();
            this.joined = true;
        },
    }
});
