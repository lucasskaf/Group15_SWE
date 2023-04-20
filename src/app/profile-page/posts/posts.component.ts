import { Component, Input, Output, EventEmitter, OnInit } from '@angular/core';
import { moviePosts } from 'src/app/user-auth/user';
import { Emmiters } from 'src/app/emitters/emmiters';
import { MovieGeneratorService } from 'src/app/services/movie-generator.service';

@Component({
  selector: 'app-posts',
  templateUrl: './posts.component.html',
  styleUrls: ['./posts.component.css']
})
export class PostsComponent implements OnInit {
  @Output() openPostAddPopup = new EventEmitter<boolean>();
  @Input() addPostPopupStatus: boolean = false;

  posts: moviePosts[] = [];

  constructor(private movieService: MovieGeneratorService) { }

  ngOnInit() {
    this.posts = Emmiters.userPosts;
  }

  onPostDelete(post: moviePosts, id: string) {
    this.movieService.removeMoviePosts(post, id).subscribe({
      next: (resp) => {
        console.log(resp);
        // remove the movie from the watchList in the backend
        const index = Emmiters.userPosts.indexOf(post);
        if (index > -1) {
          Emmiters.userPosts.splice(index, 1);
        }
      },
      error: (error) => {
        console.log(error);
      }
    });
  }
}
