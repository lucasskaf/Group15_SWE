import { Injectable } from '@angular/core';
import { PostComponent } from '../common/post/post.component';

@Injectable({
  providedIn: 'root'
})
export class PostService {
  postsMade: PostComponent[] = [];

  getPostsMade() {
    return this.postsMade;
  }

  addPost(post: PostComponent) {
    this.postsMade.push(post);
  }

  deletePost(post: PostComponent) {
    const index = this.postsMade.indexOf(post);
    if (index >= 0) {
      this.postsMade.splice(index, 1);
    }
  }
}
