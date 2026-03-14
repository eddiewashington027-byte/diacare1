def hello(name: str = "world") -> str:
    """Return a simple greeting.

    This small function exists to provide a minimal runtime target
    for CI and local tests.
    """
    return f"hello {name}"

__all__ = ["hello"]
