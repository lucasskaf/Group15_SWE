import { Component, EventEmitter, Output, Input, OnInit } from '@angular/core';
import { MovieFormComponent } from '../movie-form/movie-form.component';

@Component({
  selector: 'app-add-popup',
  templateUrl: './add-popup.component.html',
  styleUrls: ['./add-popup.component.css']
})
export class AddPopupComponent {
  @Output() isClosed = new EventEmitter();
  @Input() isOpen : boolean = false;

  ngOnInit() {
    
  }

  togglePopup() {
    this.isOpen = !this.isOpen;
    this.isClosed.emit(this.isOpen);
  }
}
