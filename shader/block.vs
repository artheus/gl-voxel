#version 330 core
in vec3 aPos;
in vec3 aNormal;

out vec3 FragPos;
out vec3 Normal;

in vec2 vertTexCoord;
out vec2 fragTexCoord;

uniform mat4 model;
uniform mat4 camera;
uniform mat4 projection;

void main()
{
    FragPos = vec3(model * vec4(aPos, 1.0));
    Normal = mat3(transpose(inverse(model))) * aNormal;

    gl_Position = projection * camera * vec4(FragPos, 1.0);
    fragTexCoord = vertTexCoord * 2;
}