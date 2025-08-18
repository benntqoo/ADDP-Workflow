// Android Kotlin code for testing
package com.example.myapp

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.compose.material3.*
import androidx.compose.runtime.*

class MainActivity : ComponentActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        
        setContent {
            MyApp()
        }
    }
}

@Composable
fun MyApp() {
    var counter by remember { mutableStateOf(0) }
    
    Button(onClick = { counter++ }) {
        Text("Clicked $counter times")
    }
}