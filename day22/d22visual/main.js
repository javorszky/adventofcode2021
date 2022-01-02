import './style.css'
import * as THREE from 'three';
import {TrackballControls} from 'three/examples/jsm/controls/TrackballControls'

// document.querySelector('#app').innerHTML = `
//   <h1>Hello Vite!</h1>
//   <a href="https://vitejs.dev/guide/features.html" target="_blank">Documentation</a>
// `
const width = window.innerWidth
const height = window.innerHeight

// Create the scene
const scene = new THREE.Scene();

// The camera
const camera = new THREE.PerspectiveCamera( 35, width / height, 1, 1000 );

const renderer = new THREE.WebGLRenderer();
renderer.setSize( window.innerWidth, window.innerHeight );
document.querySelector('#app').appendChild( renderer.domElement );

const geometry = new THREE.BoxGeometry();
const material = new THREE.MeshBasicMaterial( { color: 0x00ff00 } );
const cube = new THREE.Mesh( geometry, material );
scene.add( cube );
scene.add(new THREE.AxesHelper(55))

camera.position.z = 150;
camera.position.x = 150;
camera.position.y = 150;
camera.lookAt(0,0,0);

const controls = new TrackballControls( camera, renderer.domElement );






function animate() {
	requestAnimationFrame( animate );
	renderer.render( scene, camera );
  controls.update()
}
animate();

