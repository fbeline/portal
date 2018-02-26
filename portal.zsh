function p() {
    local output="$(portal $1)"
    outputportal output
}

function pc() {
    local output="$(portal c $(pwd) $1)"
    outputportal output
}

function outputportal() {
    if [[ -d "${output}" ]]; then
        cd "${output}"
    else
        echo "portal: directory '${@}' not found"
        echo "Try \`portal --help\` for more information."
        false
    fi
}

preexec() { portal s $(pwd) }

