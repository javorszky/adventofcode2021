import "./style.css";
import * as THREE from "three";
import { ArcballControls } from "three/examples/jsm/controls/ArcballControls.js";
import * as cubes from "./js/cubes.js";

// document.querySelector('#app').innerHTML = `
//   <h1>Hello Vite!</h1>
//   <a href="https://vitejs.dev/guide/features.html" target="_blank">Documentation</a>
// `
let INTERSECTED;
let width, height;
let camera, scene, raycaster, renderer, controls;

const pointer = new THREE.Vector2();

init();

parseCubes(cubes.default);

animate();

function init() {
  width = window.innerWidth;
  height = window.innerHeight;

  // Create the scene
  scene = new THREE.Scene();
  scene.background = new THREE.Color(0xf0f0f0);

  // The camera
  camera = new THREE.PerspectiveCamera(35, width / height, 1, 1000);
  camera.position.z = 150;
  camera.position.x = 150;
  camera.position.y = 150;
  camera.lookAt(0, 0, 0);

  // The Raycaster
  raycaster = new THREE.Raycaster();

  // The renderer
  renderer = new THREE.WebGLRenderer();
  renderer.setPixelRatio(window.devicePixelRatio);
  renderer.setSize(window.innerWidth, window.innerHeight);

  document.querySelector("#app").appendChild(renderer.domElement);

  document.addEventListener("mousemove", onPointerMove);

  // The Controls
  controls = new ArcballControls(camera, renderer.domElement, scene);

  scene.add(new THREE.AxesHelper(55));
}

function onPointerMove(event) {
  pointer.x = (event.clientX / window.innerWidth) * 2 - 1;
  pointer.y = -(event.clientY / window.innerHeight) * 2 + 1;
}

function parseCubes(cubeString) {
  for (const [key, value] of Object.entries(cubeString)) {
    const x = value.x_to - value.x_from + 1;
    const y = value.y_to - value.y_from + 1;
    const z = value.z_to - value.z_from + 1;

    const geo = new THREE.BoxGeometry(x, y, z);

    const material = new THREE.MeshStandardMaterial({
      // color: Math.random() * 0xffffff,
      emissive: Math.random() * 0xffffff,
      opacity: 0.3,
      transparent: true,
    });

    const cube = new THREE.Mesh(geo, material);
    cube.name = key;

    cube.position.x = value.x_from + x / 2;
    cube.position.y = value.y_from + y / 2;
    cube.position.z = value.z_from + z / 2;

    scene.add(cube);
  }
}

function animate() {
  requestAnimationFrame(animate);

  render();
}

function render() {
  raycaster.setFromCamera(pointer, camera);

  const intersects = raycaster.intersectObjects(scene.children, false);

  if (intersects.length > 0) {
    if (INTERSECTED != intersects[0].object) {
      if (INTERSECTED && INTERSECTED.material.type == "MeshStandardMaterial") {
        INTERSECTED.material.opacity = 0.3;
      }

      if (intersects[0].object.material.type == "MeshStandardMaterial") {
        INTERSECTED = intersects[0].object;
        INTERSECTED.material.opacity = 0.7;
      }
    }
  } else {
    if (INTERSECTED && INTERSECTED.material.type == "MeshStandardMaterial") {
      INTERSECTED.material.opacity = 0.3;
    }

    INTERSECTED = null;
  }

  renderer.render(scene, camera);
  controls.update();
}
