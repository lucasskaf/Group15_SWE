import { Component, Input, Output, EventEmitter, OnInit } from '@angular/core';
import { moviePosts } from 'src/app/user-auth/user';
import { Emmiters } from 'src/app/emitters/emmiters';

@Component({
  selector: 'app-posts',
  templateUrl: './posts.component.html',
  styleUrls: ['./posts.component.css']
})
export class PostsComponent implements OnInit {
  @Output() openPostAddPopup = new EventEmitter<boolean>();
  @Input() addPostPopupStatus: boolean = false;
  
  posts: moviePosts[] = [];

  constructor() {}
  
  ngOnInit() {
    this.posts = Emmiters.userPosts;
  }

  // ngOnInit() {
  //   this.posts = this.postService.getPostsMade();
  // }

  // onPostDelete(post: PostComponent) {
  //   this.postService.deletePost(post);
  // }

  // public onPostAddClick(): void {
  //   this.openPostAddPopup.emit(!this.addPostPopupStatus);
  //   this.addPostPopupStatus = !this.addPostPopupStatus;
  // }
}
