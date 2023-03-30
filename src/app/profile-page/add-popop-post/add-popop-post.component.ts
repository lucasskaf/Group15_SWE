import { Component, Input, Output, EventEmitter } from '@angular/core';

@Component({
  selector: 'app-add-popop-post',
  templateUrl: './add-popop-post.component.html',
  styleUrls: ['./add-popop-post.component.css']
})
export class AddPopopPostComponent {
  @Output() isClosed = new EventEmitter();
  @Input() isOpen : boolean = false;

  ngOnInit() {
    
  }

  togglePopup() {
    this.isOpen = !this.isOpen;
    this.isClosed.emit(this.isOpen);
  }
}
