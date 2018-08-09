import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { AppComponent } from './app.component';
import { GenerateGuidComponent } from './generate-guid/generate-guid.component';
import { HttpClientModule } from '@angular/common/http';
import { HomeComponent } from './home/home.component';
import { AboutComponent } from './about/about.component';
import { AppRoutes } from './app.routes';
import { GenerateQrComponent } from './generate-qr/generate-qr.component';
import { TimeConverterComponent } from './time-converter/time-converter.component';
import { EncryptionComponent } from './encryption/encryption.component';

@NgModule({
  declarations: [
    AppComponent,
    GenerateGuidComponent,
    HomeComponent,
    AboutComponent,
    GenerateQrComponent,
    TimeConverterComponent,
    EncryptionComponent
  ],
  imports: [
    BrowserModule,
    HttpClientModule,
    RouterModule.forRoot(AppRoutes, { enableTracing: true })
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
