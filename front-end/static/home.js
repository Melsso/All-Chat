// might need to add a cleaner, after logout might need to clean the html
let currentIndex = 0;

function capitalizeFirstLetterOfEachWord(str) {
    return str.split(' ')
        .map(word => word.charAt(0).toUpperCase() + word.slice(1).toLowerCase())
        .join(' ');
}

function fetchComments(postId) {
    fetch(`https://localhost:8443/add-comment?post_id=${postId}`, {
        method: 'GET',
        credentials: 'include',
        headers: {
            'Content-type': 'application/json',
        },
    })
    .then(response => {
        if (!response.ok) {
            return handleFetchError(response);
        }
        return response.json();
    })
    .then(data => {
        const commentSection = document.getElementById(`comments`);
        commentSection.innerHTML = '';

        if (data.comments && data.comments.length > 0) {
            data.comments.forEach(comment => {
                const commentField = document.createElement('article');
                commentField.classList.add('comment');
                commentField.innerHTML = `
                    <section class="comment-content">
                        <p class="comment-owner">${comment.comment_owner}:</p>
                        <p class="comment-content">${comment.content}</p>
                        <small>${new Date(comment.created_at).toLocaleString()}</small>
                    </section>
                `;
                commentSection.appendChild(commentField);
            });
        }
        else {
            commentSection.innerHTML = '<p id="nocom">No comments available.</p>';
        }
    })
    .catch(error => handleFetchError(error));
}

function fetchLike(postId) {
    fetch('https://localhost:8443/like-post', {
        method: 'POST',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ post_id: postId}),
    })
    .then(response => {
        if (!response.ok)
            return handleFetchError(response);
        return response.json();
    })
    .then(data => {
        if (data.status === 'liked')
            console.log(`Post${postId} liked successfully`);
        else
            console.log(`Failed to like Post${postId}`);
    })
    .catch(error => handleFetchError(error));
}

function AddNewComment(postId, content) {
    fetch('https://localhost:8443/add-comment', {
        method: 'POST',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ post_id: postId, content: content }),
    })
    .then(response => {
        if (!response.ok)
            return handleFetchError(response);
        return response.json();
    })
    .then(data => {
        if (data.status === 'commented') {
            console.log('Comment added successfully');
            fetchComments(postId);
        }
        else {
            alert('Failed to add comment');
        }
    })
    .catch(error => handleFetchError(error));
}

function renderPosts(posts) {
    const postsContainer = document.getElementById('posts');
    postsContainer.innerHTML = '';
    if (posts.length === 0) {
        postsContainer.innerHTML = '<p>No posts available for now</p>';
        return;
    }
    posts.forEach((post, index) => {
        const divPost = document.createElement('div');
        divPost.classList.add('post');
        divPost.setAttribute('id', `post-${post.post_id}`);
        if (index === 0) {
            divPost.classList.add('active');
        } 
        divPost.innerHTML = `
            <header>
                <h3 class="post-owner">${post.post_owner}:</h3>
                <small>${new Date(post.created_at).toLocaleString()}</small>
            </header>

            <section class="post-content">
                <p>${post.content}</p>
            </section>
        `
        postsContainer.appendChild(divPost);
    });
}

function renderFriends(friends) {
    const friendsContainer = document.getElementById('friend-list');
    friendsContainer.innerHTML = '';
    if (friends && friends.length > 0) {
        friends.forEach(friend => {
            const li = document.createElement('li');
            li.className = 'friend';
            li.textContent = friend.first_name + " " + friend.last_name;

            const deleteButton = document.createElement('button');
            const messageButton = document.createElement('button');

            deleteButton.textContent = '✖';
            deleteButton.className = 'delete-button';
            deleteButton.title = 'Delete Friend';
            deleteButton.setAttribute('data-friend-id', friend.user_id);
            messageButton.textContent = '✉';
            messageButton.className = 'message-button';
            messageButton.title = 'Message Friend';
            messageButton.setAttribute('data-friend-id', friend.user_id);
        
            li.appendChild(deleteButton);
            li.appendChild(messageButton);
            friendsContainer.appendChild(li);
        });
    } else {
        friendsContainer.innerHTML = '<p id="nofriend">No Friends Yet.</p>'
    }
}

function renderInviteList(invites) {
    const inviteListContainer = document.getElementById('invite-list');
    inviteListContainer.innerHTML = '';
    if (invites && invites.length > 0) {
        invites.forEach(invite => {
            const li = document.createElement('li');
            li.className = 'invite';
            li.textContent = invite.first_name + " " + invite.last_name;

            const acceptButton = document.createElement('button');
            const refuseButton = document.createElement('button');

            acceptButton.textContent = '✔';
            acceptButton.className = 'accept-button';
            acceptButton.title = 'Accept Friend Request';
            acceptButton.setAttribute('data-friend-id', invite.user_id);
            refuseButton.textContent = '✖';
            refuseButton.className = 'refuse-button';
            refuseButton.title = 'Refuse Friend Request';
            refuseButton.setAttribute('data-friend-id', invite.user_id);
        
            li.appendChild(acceptButton);
            li.appendChild(refuseButton);
            inviteListContainer.appendChild(li);
        });
    } else {
        inviteListContainer.innerHTML = '<p id="nofriend">No Friend Requests.</p>'
    }
}

function fetchContents() {
    fetch('https://localhost:8443/home', {
        method: 'GET',
        credentials: 'include',
        headers: {
            'Content-type': 'application/json',
        },
    })
    .then(response => {
        if (!response.ok) {
            handleFetchError(response);
        }
        return response.json();
    })
    .then(data => {
        if (data.posts) {// CHECK REFUSE FRIEND REQUEST BUTTON

            renderPosts(data.posts);
            renderFriends(data.friends);
            renderInviteList(data.invite);
        } else {
            console.error('No posts/friends data received:', data);
        }
    })
    .catch(error => handleFetchError(error));
}

function handleFetchError(error) {
    console.error("Following error occured: ", error.message);
}

function updateCurrent(direction) {
    const carouselItems = document.querySelectorAll('.post');
    carouselItems[currentIndex].classList.remove('active');

    currentIndex = (currentIndex + direction + carouselItems.length) % carouselItems.length;
    carouselItems[currentIndex].classList.add('active');
}

function getActivePostId() {
    const activePost = document.querySelector('.post.active');
    if (activePost) {
        return activePost.id.split('-')[1];
    } else {
        console.error('No active post found.');
        return null;
    }
}

// Function to create a new post
function createPost(content) {
    fetch('https://localhost:8443/create-post', {
        method: 'POST',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ content: content }),
    })
    .then(response => {
        if (!response.ok) {
            return handleFetchError(response);
        }
        return response.json();
    })
    .then(newPost => {
        console.log('New post created:', newPost);
        fetchContents();
    })
    .catch(error => handleFetchError(error));
}

// Function to logout the user and invalidate the cookie
function logoutUser() {
    fetch('https://localhost:8443/logout', {
        method: 'POST',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
        }
    }).then(response => {
        if (response.ok) {
            document.cookie = "auth" + '=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/; domain=example.com;';
            window.location.href = '/login.html';
        }
    });
}    

// Function that handles all the side buttons and their event listeners
function SidebarButtonsHandlers() {    
    document.addEventListener('click', function(event) {
        if (event.target.className == 'delete-button') {
            personId = event.target.getAttribute('data-friend-id');        
            deleteFriend(personId);
        
        } else if (event.target.className == 'message-button') {
            personId = event.target.getAttribute('data-friend-id');
            const parentElement = event.target.parentElement;
            const fullText = parentElement.textContent.trim();
            personName = fullText.split('✖')[0].trim();
            openMessageWindow(personId, personName);
        
        } else if (event.target.className == 'search-button') {
            event.preventDefault();
            const unameText = document.getElementById('search-input');
            const uname = unameText.value;
            if (uname) {
                lookupUser(uname);
                unameText.value = '';
            } else {
                alert('Please enter a name');
            }
        
        } else if (event.target.className == 'accept-button') {
            personId = event.target.getAttribute('data-friend-id');
            acceptFriend(personId, "y");

        } else if (event.target.className == 'refuse-button') {
            personId = event.target.getAttribute('data-friend-id');
            acceptFriend(personId, "n");
        
        } else if (event.target.className == 'add-ser-button') {
            personId = event.target.getAttribute('data-result-id');
            addFriend(personId);
            const btn = event.target;
            btn.textContent = 'SENT!';
            btn.disabled = true;
            // fetchContents();

        } else if (event.target.id == 'submit-post') {
            const content = document.getElementById('post-post').value;
            if (!content) {
                alert('Please enter some content for your post.');
                return ;
            }
            const element = document.getElementById('post-post');
            element.value = '';
            createPost(content);
        
        } else if (event.target.id == 'dark-mode') {
            document.body.classList.toggle('invert-mode');
            if (document.body.classList.contains('invert-mode')) {
                event.target.textContent = 'Default Mode';
            } else {
                event.target.textContent = 'Invert Colors';
            }
        
        } else if (event.target.id == 'logout-link') {
           logoutUser();
        }
    })
}

// Function that handles all the carousel buttons and their event listeners
function CarouselButtonsHandlers() {
    const nextButton = document.getElementById('carousel-control-next');
    const prevButton = document.getElementById('carousel-control-prev');
    const likeButton = document.getElementById('like-button');
    const commentsButton = document.getElementById('comment-button');
    const submitButton = document.getElementById('submit-comment-button');

    likeButton.addEventListener('click', () => {
        const postId = getActivePostId();
        if (postId) fetchLike(postId);
    });
    
    commentsButton.addEventListener('click', () => {
        const postId = getActivePostId();
        if (postId) {
            const commentSection = document.getElementById(`comments`);
            if (commentSection.classList.contains('displayed')) {
                commentSection.innerHTML = '';
                commentSection.classList.remove('displayed');
            } else {
                fetchComments(postId);
                commentSection.classList.add('displayed');
            }
        }
    });

    submitButton.addEventListener('click', () => {
        const postId = getActivePostId();
        const contentText = document.getElementById('com-content');
        const content = contentText.value;
        if (postId && content) {
            AddNewComment(postId, content);
            contentText.value = '';
        } else {
            alert('Please enter some content for your comment.');
        }
    });

    nextButton.addEventListener('click', () => updateCurrent(1));
    prevButton.addEventListener('click', () => updateCurrent(-1));
}

// Function that fetches the messages from backend
function openMessageWindow(friendId, friendName) {
    fetch(`https://localhost:8443/messages?friend_id=${friendId}`, {
        method: 'GET',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
        },
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to get messages');
        }
        return response.json();
    })
    .then(data => {
        if (data.messages) {
            showMessageWindow(data.conversation_id, data.messages, friendId, friendName);
        }
        else {
            console.error('No messages data received:', data);
        }
    })
    .catch(error => handleFetchError(error));
}

// Function to send a message to someone
function sendMessage(conversationId, content) {
    fetch('https://localhost:8443/messages', {
        method: 'POST',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ conversation_id: conversationId, content: content }),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to send message');
        }
        return response.json();
    })
    .then(data => {
        if (data.status === 'sent') {
            console.log('Message sent successfully');
        } else {
            console.error('Failed to send message: ', error.message);
        }
    })
    .catch(error => handleFetchError(error));
}

// Function to create a friend request
function addFriend(friendId) {
    fetch('https://localhost:8443/add-friend', {
        method: 'POST',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ friend_id: friendId }),
    })
    .then(response => {
        if (!response.ok) {
            alert('Failed to add friend...');
            throw new Error('Failed to add friend');
        }
        return response.json();
    })
    .then(data => {
        // alert('', data.message);
        fetchContents();
    })
    .catch(error => console.error('Error adding friend:', error.message));
}

// Fuction to accept a friend request
function acceptFriend(friendId, choice) {
    fetch('https://localhost:8443/accept-friend', {
        method: 'POST',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ friend_id: friendId, action: choice}),
    })
    .then(response => {
        if (!response.ok) {
            alert('Failed to process friend request...');
            throw new Error('Failed to accept friend');
        }
        return response.json();
    })
    .then(data => {
        if (choice == "y") {
            alert('Friend Accepted.')
        }
        else {
            alert('Friend Denied');
        }
        fetchContents();
    })
    .catch(error => console.error('Error accepting friend:', error.message));
}

// Function to handle deleting a friend
function deleteFriend(friendId) {
    fetch('https://localhost:8443/delete-friend', {
        method: 'POST',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ friend_id: friendId }),
    })
    .then(response => {
        if (!response.ok) {
            alert('Failed to delete friend');
            throw new Error('Failed to delete friend');
        }
        return response.json();
    })
    .then(data => {
        alert('Friend deleted.');
        fetchContents();
    })
    .catch(error => console.error('Error deleting friend:', error.message));
}

// Function to handle user look up, will need to rename the variables, username does not exist, it is missleading
function lookupUser(username) {
    fetch(`https://localhost:8443/lookup-user`, {
        method: 'POST',
        credentials: 'include',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({user_name: username}),
    })
    .then(response => {
        if (!response.ok) {
            console.log('failure: ', username);
            throw new Error('Failed to lookup user');
        }
        return response.json();
    })
    .then(data => {
        searchResults(data.user_list);
    })
    .catch(error => console.error('Error looking up user:', error.message));
}

function searchResults(results) {
    const resContainer = document.getElementById('search-result');
    resContainer.innerHTML = '';

    if (results && results.length > 0) {
        results.forEach(result => {
            const li = document.createElement('li');
            li.className = "search-item";
            li.textContent = result.first_name + " " + result.last_name;

            const addButton = document.createElement('button');
            addButton.textContent = 'INVITE';
            addButton.className = 'add-ser-button';
            addButton.title = 'Send Friend Request';
            addButton.setAttribute('data-result-id', result.user_id);

            li.appendChild(addButton);
            resContainer.appendChild(li);
        });
    } else {
        resContainer.innerHTML = '<p id="nores">No Results.</p>';
    }
}


function showMessageWindow(conversationId, messages, friendId, friendName) {
    // Check if a message window already exists and remove it
    const existingWindow = document.querySelector('.message-window');
    if (existingWindow) {
        const existingConversationId = existingWindow.dataset.conversationId;
        if (existingConversationId === conversationId.toString()) {
            // If the same conversation ID, remove the window (toggle behavior)
            existingWindow.remove();
            return;
        } else {
            // If a different conversation ID, remove the existing window
            existingWindow.remove();
        }
    }

    // Create new message window
    const messageWindow = document.createElement('div');
    messageWindow.className = 'message-window';
    messageWindow.dataset.conversationId = conversationId;

    const header = document.createElement('header');
    header.textContent = `${friendName}`;
    header.style.color = '#00008B';
    header.style.textAlign = 'center'; 
    header.style.textTransform = 'uppercase';
    messageWindow.appendChild(header);

    const messageContainer = document.createElement('div');
    messageContainer.className = 'message-container';
    messages.forEach(message => {
        const messageElement = document.createElement('p');
        
        const senderName = document.createElement('span');
        senderName.style.color = '#00008B';
        senderName.style.fontWeight = 'bold';
        if (message.sender_id == friendId) {
            senderName.textContent = `${capitalizeFirstLetterOfEachWord(friendName)}: `;
        }
        else {
            senderName.textContent = `You: `;
        }
        messageElement.appendChild(senderName);
        messageElement.appendChild(document.createTextNode(message.content));

        messageElement.style.color = 'black';
        messageContainer.appendChild(messageElement);
    });
    messageWindow.appendChild(messageContainer);

    const textarea = document.createElement('textarea');
    textarea.placeholder = 'Type your message...';
    messageWindow.appendChild(textarea);

    const sendButton = document.createElement('button');
    sendButton.textContent = 'Send';
    sendButton.addEventListener('click', function() {
        const content = textarea.value;
        if (!content) {
            alert('Please enter a message');
            return;
        }
        sendMessage(conversationId, content);
        openMessageWindow(friendId, friendName);
    });
    messageWindow.appendChild(sendButton);

    const messageWindowContainer = document.getElementById('message-window-container');
    messageWindowContainer.appendChild(messageWindow);
}



document.addEventListener('DOMContentLoaded', function() {

    fetchContents();
    CarouselButtonsHandlers();
    SidebarButtonsHandlers();
});