import { Component, Input, Output, EventEmitter } from '@angular/core';
import { PostService } from '../../services/post-service';
import { PostComponent } from 'src/app/common/post/post.component';

@Component({
  selector: 'app-posts',
  templateUrl: './posts.component.html',
  styleUrls: ['./posts.component.css']
})
export class PostsComponent {
  @Output() openPostAddPopup = new EventEmitter<boolean>();
  @Input() addPostPopupStatus: boolean = false;
  
  posts: PostComponent[] = [];

  constructor(private postService: PostService) {}

  ngOnInit() {
    this.posts = this.postService.getPostsMade();
  }

  onPostDelete(post: PostComponent) {
    this.postService.deletePost(post);
  }

  public onPostAddClick(): void {
    this.openPostAddPopup.emit(!this.addPostPopupStatus);
    this.addPostPopupStatus = !this.addPostPopupStatus;
  }
}
