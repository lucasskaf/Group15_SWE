import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatToolbarModule} from '@angular/material/toolbar';
import { MatButtonModule } from '@angular/material/button';
import { NavbarComponent } from './navbar/navbar.component';
import { SidebarComponent } from './sidebar/sidebar.component';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatListModule } from '@angular/material/list';
import { MatIconModule } from '@angular/material/icon';
import { MainComponent } from './main/main.component';
import { WatchedComponent } from './watched/watched.component';
import { PostsComponent } from './posts/posts.component';
import { NavBarComponent } from '../nav-bar/nav-bar.component';

@NgModule({
  declarations: [
    SidebarComponent,
    MainComponent,
    WatchedComponent,
    PostsComponent,
    NavbarComponent
  ],
  imports: [
    CommonModule,
    MatToolbarModule,
    MatButtonModule,
    MatSidenavModule,
    MatListModule,
    MatIconModule
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
