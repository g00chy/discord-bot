import React, { Component } from 'react'

class CommentList extends Component {
  constructor() {
    super();
    this.state = {
      comments: [
        {
          id: 1,
          comment: 'comment 1'
        },
        {
          id: 2,
          comment: 'comment 2'
        },
        {
          id: 3,
          comment: 'comment 3'
        },
        {
          id: 4,
          comment: 'comment 4'
        }
      ],
      hasError: false,
      isLoading: false
    }
  }
  render() {
    if (this.state.hasError) {
      return <p>error</p>;
    }
    if (this.state.isLoading) {
      return <p>loading . . . </p>;
    }
    return (
      <ul>
        {this.state.comments.map((item) => (
          <li key={item.id}>
            {item.comment}
          </li>
        ))}
      </ul>
    )
  }
}

export default CommentList;
