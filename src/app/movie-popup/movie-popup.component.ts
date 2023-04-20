import { Component, EventEmitter, Input, Output } from '@angular/core';
import { Movie, moviePosts } from '../user-auth/user';
import { OnInit } from '@angular/core';
import { Emmiters } from '../emitters/emmiters';
import { MovieGeneratorService } from '../services/movie-generator.service';
import { FormBuilder, Validators } from '@angular/forms';

@Component({
  selector: 'app-movie-popup',
  templateUrl: './movie-popup.component.html',
  styleUrls: ['./movie-popup.component.css']
})
export class MoviePopupComponent implements OnInit {
  postForm
  username: string = ""
  isPostOpen: boolean = false
  didSendReview: boolean = false
  didFoundPosts: boolean = false
  moviePostsEmmiter: EventEmitter<moviePosts[]> = new EventEmitter()
  moviePosts: moviePosts[] = []
  moviePostsValues: moviePosts[] = []
  isAuthenticated: boolean = false
  isMovieWatched: boolean = false
  isPopupOpen: boolean = false
  generatedMovie: Movie = {id: 0, vote_average: 0}

  constructor(private movieGeneratorService : MovieGeneratorService,
    private formBuilder: FormBuilder) {
    }

  ngOnInit(): void {
    Emmiters.authEmmiter.subscribe(
      {
        next: (auth : boolean) => {
          this.isAuthenticated = auth
        }
      }
    )

    Emmiters.userData.subscribe(
      {
        next: (username: string) => {
          this.username = username
        }
      }
    )

    Emmiters.generatedMovie.subscribe(
      {
        next: (movie) => {
          this.generatedMovie = movie
        },
        error: (err) => {
          console.log(`Was not able to get movie`)
        }
      }
    )

    Emmiters.isPopupOpen.subscribe(
      {
        next: (status) => {
          this.isPopupOpen = status

          if(status){
            console.log(this.generatedMovie.id, this.generatedMovie.title)
            this.generateMoviePosts(this.generatedMovie.id.toString())
          }
        }
      }
    )

    Emmiters.isMovieWatched.subscribe(
      {
        next: (status) => {
          this.isMovieWatched = status
        }
      }
    )

    this.moviePostsEmmiter.subscribe({
      next: (newMoviePost) => {
        this.moviePostsValues = newMoviePost
      },
      error: (err) => {
        console.log('IN GETTING POST VALUES', err)
      }
    })

    this.postForm = this.formBuilder.group({
      movie_id: this.generatedMovie.id,
      username: this.username,
      title: this.formBuilder.control(''),
      body: this.formBuilder.control('', Validators.required)
    })
  }

  closeMovie() {
    console.log('CLICKED ON CLOSE')
    Emmiters.isPopupOpen.emit(false)
    Emmiters.isMovieWatched.emit(false)
    this.isMovieWatched = false
  }

  toogleAddPost(){
    this.isPostOpen = !this.isPostOpen
  }

  onSubmitPost(post: moviePosts){
    post.movie_id = this.generatedMovie.id.toString()
    post.username = this.username
    console.log('THIS IS THE POST:')
    console.log(post)

    this.movieGeneratorService.addMoviePosts(post).subscribe(
      {
        next: (resp) => {
          this.moviePosts.push(resp)
          this.moviePostsEmmiter.emit(this.moviePosts)
          this.didFoundPosts = true
          this.didSendReview = true
          this.isPostOpen = false
        },
        error: (err) => {
          console.log('NOT ADD POST')
        }
      }
    )
  }

  addToWatchlist(movie: Movie) {
    this.movieGeneratorService.addToWatchList(movie).subscribe(
      {
        next: (resp) => {
          console.log(resp)
          this.isMovieWatched = true
          Emmiters.isMovieWatched.emit(true)
          console.log(`BEFORE ADDED: ${Emmiters.watchList.length}`)

          Emmiters.watchList.push(movie)
        },
        error: (err) => {
          console.log(err)
        }
      }
    )
  }

  generateMoviePosts(movieID: string | undefined){
    this.movieGeneratorService.getMoviePosts(movieID).subscribe(
      {
        next: (resp) => {
          console.log(`POSTS: ${resp}`)
          this.moviePosts = resp
          this.moviePostsEmmiter.emit(this.moviePosts)
          this.didFoundPosts = true
        },
        error: (err) => {
          console.log('ERROR IN GETTING POSTS')
          if(err.status == 404){
            this.didFoundPosts = false
          }
        }
      }
    )
  }
}
