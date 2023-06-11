import Home from "../components/Home.vue";
import File from "../components/File.vue"
import Note from "../components/Note.vue"
import Upload from "../components/Upload.vue"
import NotFound from "../components/NotFound.vue"
import { createRouter, createWebHistory } from "vue-router";


const routes = [
    {
        path: "/",
        name: "home",
        component: Home,
    },
    {
        path: "/up/:type",
        name: "upload",
        component: Upload,
        props: true
    },
    {
        path: "/f/:shareCode",
        name: "file",
        component: File,
        props: true, // can accept route parameter as props
    },
    {
        path: "/n/:shareCode",
        name: "note",
        component: Note,
        props: true, // can accept route parameter as props
    },
    // catchall 404
    {
        path: "/:catchAll(.*)", // regex, catch all that does not match any of the above
        name: "Not Found",
        component: NotFound,
    },
];

const router = createRouter({
    // 4. 内部提供了 history 模式的实现。为了简单起见，我们在这里使用 hash 模式。
    history: createWebHistory(),
    routes, // `routes: routes` 的缩写
})

export default router;