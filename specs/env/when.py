import requests as r
import env
import urllib.parse as url


def view_dataset(ds_id):
    return _get_page("view_dataset",  env.urls.view_dataset(ds_id))


def search_datasets(query):
    return _get_page(f"search datasets for {query}", env.urls.search_datasets(query))


def view_project(project_id):
    return _get_page("view project", env.urls.view_project(project_id))


def list_datasets():
    return _get_page("list datasets", env.urls.list_datasets())


def list_projects():
    return _get_page("list projects", env.urls.list_projects())


def _get_page(text, uri):
    def _(c: r.Session, base: str):
        full = url.urljoin(base, uri)
        return c.get(full)

    return env.When(action=_, text=text)


