import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatToolbarModule} from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';
import { NavbarComponent } from './navbar/navbar.component';
import { ReactiveFormsModule } from '@angular/forms';
import { SidebarComponent } from './sidebar/sidebar.component';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatListModule } from '@angular/material/list';
import { MatIconModule } from '@angular/material/icon';
import { MainComponent } from './main/main.component';
import { WatchedComponent } from './watched/watched.component';
import { PostsComponent } from './posts/posts.component';
import { MovieFormComponent } from './movie-form/movie-form.component';
import { MovieComponent } from '../common/movie/movie.component';
import { AddMoviePopupComponent } from './add-popup-movie/add-popup-movie.component';
import { PostComponent } from '../common/post/post.component';
import { AddPopopPostComponent } from './add-popop-post/add-popop-post.component';
import { PostFormComponent } from './post-form/post-form.component';
import { routing } from '../app-routing.module';

@NgModule({
  declarations: [
    SidebarComponent,
    MainComponent,
    WatchedComponent,
    PostsComponent,
    NavbarComponent,
    MovieFormComponent,
    MovieComponent,
    AddMoviePopupComponent,
    PostComponent,
    AddPopopPostComponent,
    PostFormComponent
  ],
  imports: [
    CommonModule,
    MatToolbarModule,
    MatButtonModule,
    MatSidenavModule,
    MatListModule,
    MatIconModule,
    ReactiveFormsModule,
    routing
  ],
  exports: [
    MatToolbarModule,
    MatButtonModule,
    MatSidenavModule,
    MatListModule,
    MatIconModule,
    MainComponent,
    NavbarComponent,
    SidebarComponent
  ]
})
export class ProfilePageModule { }