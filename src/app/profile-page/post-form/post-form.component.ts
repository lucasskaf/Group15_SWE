import { Component } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms'
import { PostService } from 'src/app/services/post-service';
import { PostComponent } from 'src/app/common/post/post.component';

@Component({
  selector: 'app-post-form',
  templateUrl: './post-form.component.html',
  styleUrls: ['./post-form.component.css']
})
export class PostFormComponent {
  form!: FormGroup;

  constructor(private postService: PostService) {}

  ngOnInit() {
    this.form = new FormGroup({
      content: new FormControl('', Validators.required )
    });
  }

  onSubmit(post: PostComponent) {
    // console.log(post);
    this.postService.addPost(post);
    this.form.reset();
  }
}
